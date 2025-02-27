/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	logger "github.com/sirupsen/logrus"
	k8error "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	k8client "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/source"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	"github.com/wso2/apk/adapter/pkg/logging"
	"github.com/wso2/apk/adapter/pkg/utils/envutils"
	"github.com/wso2/apk/adapter/pkg/utils/stringutils"
	cache "github.com/wso2/apk/common-controller/internal/cache"
	"github.com/wso2/apk/common-controller/internal/config"
	loggers "github.com/wso2/apk/common-controller/internal/loggers"
	dpv1alpha1 "github.com/wso2/apk/common-controller/internal/operator/api/v1alpha1"
	constants "github.com/wso2/apk/common-controller/internal/operator/constant"
	xds "github.com/wso2/apk/common-controller/internal/xds"
)

// RateLimitPolicyReconciler reconciles a RateLimitPolicy object
type RateLimitPolicyReconciler struct {
	client client.Client
	ods    *cache.RatelimitDataStore
	Scheme *runtime.Scheme
}

const (
	// apiRateLimitIndex Index for API level ratelimits
	apiRateLimitIndex = "apiRateLimitIndex"
	// apiRateLimitResourceIndex Index for resource level ratelimits
	httprouteRateLimitIndex = "httprouteRateLimitIndex"
)

// NewratelimitController creates a new ratelimitcontroller instance.
func NewratelimitController(mgr manager.Manager, ratelimitStore *cache.RatelimitDataStore) error {
	ratelimitReconsiler := &RateLimitPolicyReconciler{
		client: mgr.GetClient(),
		ods:    ratelimitStore,
	}
	ctx := context.Background()
	if err := addIndexes(ctx, mgr); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2612, err))
		return err
	}

	c, err := controller.New(constants.RatelimitController, mgr, controller.Options{Reconciler: ratelimitReconsiler})
	if err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(3111, err.Error()))
		return err
	}

	if err := c.Watch(&source.Kind{Type: &dpv1alpha1.API{}},
		handler.EnqueueRequestsFromMapFunc(ratelimitReconsiler.getRatelimitForAPI)); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2611, err))
		return err
	}

	if err := c.Watch(&source.Kind{Type: &gwapiv1b1.HTTPRoute{}},
		handler.EnqueueRequestsFromMapFunc(ratelimitReconsiler.getRatelimitForHTTPRoute)); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2611, err))
		return err
	}

	if err := c.Watch(&source.Kind{Type: &dpv1alpha1.RateLimitPolicy{}}, &handler.EnqueueRequestForObject{}); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(3112, err.Error()))
		return err
	}

	loggers.LoggerAPKOperator.Debug("RatelimitPolicy Controller successfully started. Watching RatelimitPolicy Objects...")
	return nil
}

//+kubebuilder:rbac:groups=dp.wso2.com,resources=ratelimitpolicies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dp.wso2.com,resources=ratelimitpolicies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dp.wso2.com,resources=ratelimitpolicies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RateLimitPolicy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (ratelimitReconsiler *RateLimitPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// Check whether the Ratelimit CR exist, if not consider as a DELETE event.
	loggers.LoggerAPKOperator.Infof("Reconciling ratelimit...")
	conf := config.ReadConfigs()
	ratelimitKey := req.NamespacedName
	var ratelimitPolicy dpv1alpha1.RateLimitPolicy

	// Check k8s RatelimitPolicy Availbility
	if err := ratelimitReconsiler.client.Get(ctx, ratelimitKey, &ratelimitPolicy); err != nil {
		resolveRateLimitAPIPolicy, found := ratelimitReconsiler.ods.GetCachedRatelimitPolicy(req.NamespacedName)
		// If availble in cache Delete cache and xds
		if found && k8error.IsNotFound(err) {
			ratelimitReconsiler.ods.DeleteCachedRatelimitPolicy(req.NamespacedName)
			logger.Info("delete api ratelimit")
			logger.Info("resolveRateLimitAPIPolicy", resolveRateLimitAPIPolicy)
			xds.DeleteAPILevelRateLimitPolicies(resolveRateLimitAPIPolicy)
			xds.UpdateRateLimiterPolicies(conf.CommonController.Server.Label)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, nil
	}
	logger.Info("add api ratelimit")
	var vhost, resolveRatelimit = ratelimitReconsiler.marshelRateLimit(ctx, ratelimitKey, ratelimitPolicy)
	logger.Info("resolveRatelimitxxxx", resolveRatelimit)
	if vhost == nil {
		logger.Info("vhost is nil hence API deleted")
		return ctrl.Result{}, nil
	}
	ratelimitReconsiler.ods.AddorUpdateRatelimitToStore(ratelimitKey, resolveRatelimit)
	xds.UpdateRateLimitXDSCache(vhost, resolveRatelimit)
	xds.UpdateRateLimiterPolicies(conf.CommonController.Server.Label)

	return ctrl.Result{}, nil
}

func (ratelimitReconsiler *RateLimitPolicyReconciler) getRatelimitForAPI(obj k8client.Object) []reconcile.Request {
	ctx := context.Background()
	api, ok := obj.(*dpv1alpha1.API)
	if !ok {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2624, api))
		return []reconcile.Request{}
	}

	requests := []reconcile.Request{}

	ratelimitPolicyList := &dpv1alpha1.RateLimitPolicyList{}
	if err := ratelimitReconsiler.client.List(ctx, ratelimitPolicyList, &k8client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(apiRateLimitIndex, NamespacedName(api).String()),
	}); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2649, NamespacedName(api).String()))
		return []reconcile.Request{}
	}

	for _, ratelimitPolicy := range ratelimitPolicyList.Items {
		requests = append(requests, ratelimitReconsiler.AddRatelimitRequest(&ratelimitPolicy)...)
	}

	return requests
}

func (ratelimitReconsiler *RateLimitPolicyReconciler) getRatelimitForHTTPRoute(obj k8client.Object) []reconcile.Request {
	ctx := context.Background()
	httpRoute, ok := obj.(*gwapiv1b1.HTTPRoute)
	if !ok {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2624, httpRoute))
		return []reconcile.Request{}
	}

	requests := []reconcile.Request{}

	ratelimitPolicyList := &dpv1alpha1.RateLimitPolicyList{}
	if err := ratelimitReconsiler.client.List(ctx, ratelimitPolicyList, &k8client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(httprouteRateLimitIndex, NamespacedName(httpRoute).String()),
	}); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2649, NamespacedName(httpRoute).String()))
		return []reconcile.Request{}
	}

	for _, ratelimitPolicy := range ratelimitPolicyList.Items {
		requests = append(requests, ratelimitReconsiler.AddRatelimitRequest(&ratelimitPolicy)...)
	}

	return requests
}

// AddRatelimitRequest adds a request to reconcile for the given ratelimit policy
func (ratelimitReconsiler *RateLimitPolicyReconciler) AddRatelimitRequest(obj k8client.Object) []reconcile.Request {
	ratelimitPolicy, ok := obj.(*dpv1alpha1.RateLimitPolicy)
	if !ok {
		loggers.LoggerAPKOperator.ErrorC(logging.GetErrorByCode(2624, ratelimitPolicy))
		return nil
	}

	if !(ratelimitPolicy.Spec.TargetRef.Kind == constants.KindAPI) {
		return nil
	}

	return []reconcile.Request{{
		NamespacedName: types.NamespacedName{
			Name: string(ratelimitPolicy.Name),
			Namespace: GetNamespace(
				(*gwapiv1b1.Namespace)(ratelimitPolicy.Spec.TargetRef.Namespace), ratelimitPolicy.Namespace),
		},
	}}
}

func (ratelimitReconsiler *RateLimitPolicyReconciler) marshelRateLimit(ctx context.Context, ratelimitKey types.NamespacedName,
	ratelimitPolicy dpv1alpha1.RateLimitPolicy) ([]string, dpv1alpha1.ResolveRateLimitAPIPolicy) {
	var api dpv1alpha1.API
	var vhost []string
	var resolveResourceList []dpv1alpha1.ResolveResource
	var resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy

	// API Level Rate limit policy
	if ratelimitPolicy.Spec.TargetRef.Kind == constants.KindAPI {
		if err := ratelimitReconsiler.client.Get(ctx, types.NamespacedName{
			Namespace: ratelimitKey.Namespace,
			Name:      string(ratelimitPolicy.Spec.TargetRef.Name)},
			&api); err != nil {
			logger.Info("errorss", err)
			return nil, resolveRatelimit
		}

		var organization = api.Spec.Organization
		var context = api.Spec.Context
		var httpRoute gwapiv1b1.HTTPRoute
		if len(api.Spec.Production) > 0 {
			for _, ref := range api.Spec.Production[0].HTTPRouteRefs {
				if ref != "" {
					if err := ratelimitReconsiler.client.Get(ctx, types.NamespacedName{
						Namespace: ratelimitKey.Namespace,
						Name:      ref},
						&httpRoute); err != nil {
						logger.Info("error", err)
					}
					for _, hostName := range httpRoute.Spec.Hostnames {
						vhost = append(vhost, string(hostName))
					}
				}
			}
		}
		if len(api.Spec.Sandbox) > 0 {
			for _, ref := range api.Spec.Sandbox[0].HTTPRouteRefs {
				if ref != "" {
					ratelimitReconsiler.client.Get(ctx, types.NamespacedName{
						Namespace: ratelimitKey.Namespace,
						Name:      ref},
						&httpRoute)
					for _, hostName := range httpRoute.Spec.Hostnames {
						vhost = append(vhost, string(hostName))
					}
				}
			}
		}
		resolveRatelimit.API.RequestsPerUnit = ratelimitPolicy.Spec.Default.API.RateLimit.RequestsPerUnit
		resolveRatelimit.API.Unit = ratelimitPolicy.Spec.Default.API.RateLimit.Unit
		resolveRatelimit.Vhost = vhost
		resolveRatelimit.Organization = organization
		resolveRatelimit.Context = context
		resolveRatelimit.UUID = string(api.ObjectMeta.UID)
	}

	// Resource Level Rate limit policy
	if ratelimitPolicy.Spec.TargetRef.Kind == constants.KindResource {
		ratelimitReconsiler.client.Get(ctx, types.NamespacedName{
			Namespace: ratelimitKey.Namespace,
			Name:      string(ratelimitPolicy.Spec.TargetRef.Name)},
			&api)
		var organization = api.Spec.Organization
		var context = api.Spec.Context
		var httpRoute gwapiv1b1.HTTPRoute
		if len(api.Spec.Production) > 0 {
			for _, ref := range api.Spec.Production[0].HTTPRouteRefs {
				if ref != "" {
					ratelimitReconsiler.client.Get(ctx, types.NamespacedName{
						Namespace: ratelimitKey.Namespace,
						Name:      ref},
						&httpRoute)
					for _, rule := range httpRoute.Spec.Rules {
						for _, filter := range rule.Filters {
							if filter.ExtensionRef != nil {
								if filter.ExtensionRef.Kind == constants.KindRateLimitPolicy {
									var resolveResource dpv1alpha1.ResolveResource
									resolveResource.Path = *rule.Matches[0].Path.Value
									resolveResource.Method = string(*rule.Matches[0].Method)
									resolveResource.PathMatchType = *rule.Matches[0].Path.Type
									resolveResourceList = append(resolveResourceList, resolveResource)
								}
								for _, hostName := range httpRoute.Spec.Hostnames {
									vhost = append(vhost, string(hostName))
								}
							}
						}

					}
				}
			}
		}
		if len(api.Spec.Sandbox) > 0 {
			for _, ref := range api.Spec.Sandbox[0].HTTPRouteRefs {
				if ref != "" {
					ratelimitReconsiler.client.Get(ctx, types.NamespacedName{
						Namespace: ratelimitKey.Namespace,
						Name:      ref},
						&httpRoute)
					for _, rule := range httpRoute.Spec.Rules {
						for _, filter := range rule.Filters {
							if filter.ExtensionRef != nil {
								if filter.ExtensionRef.Kind == constants.KindRateLimitPolicy {
									var resolveResource dpv1alpha1.ResolveResource
									resolveResource.Path = *rule.Matches[0].Path.Value
									resolveResource.Method = string(*rule.Matches[0].Method)
									resolveResource.PathMatchType = *rule.Matches[0].Path.Type
									resolveResourceList = append(resolveResourceList, resolveResource)
								}
								for _, hostName := range httpRoute.Spec.Hostnames {
									vhost = append(vhost, string(hostName))
								}
							}
						}

					}
				}
			}
		}
		resolveRatelimit.Organization = organization
		resolveRatelimit.Context = context
		resolveRatelimit.UUID = string(api.ObjectMeta.UID)
		resolveRatelimit.Vhost = vhost
		resolveRatelimit.Resources = resolveResourceList
		resolveRatelimit.API.RequestsPerUnit = ratelimitPolicy.Spec.Default.API.RateLimit.RequestsPerUnit
		resolveRatelimit.API.Unit = ratelimitPolicy.Spec.Default.API.RateLimit.Unit
	}
	return vhost, resolveRatelimit
}

func addIndexes(ctx context.Context, mgr manager.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(ctx, &dpv1alpha1.RateLimitPolicy{}, httprouteRateLimitIndex,
		func(rawObj k8client.Object) []string {
			ratelimitPolicy := rawObj.(*dpv1alpha1.RateLimitPolicy)
			var apis []string
			if ratelimitPolicy.Spec.TargetRef.Kind == constants.KindResource {
				apis = append(apis,
					types.NamespacedName{
						Namespace: GetNamespace(
							(*gwapiv1b1.Namespace)(ratelimitPolicy.Spec.TargetRef.Namespace),
							ratelimitPolicy.Namespace),
						Name: string(ratelimitPolicy.Spec.TargetRef.Name),
					}.String())
			}
			return apis
		}); err != nil {
		return err
	}

	// ratelimite policy to API indexer
	err := mgr.GetFieldIndexer().IndexField(ctx, &dpv1alpha1.RateLimitPolicy{}, apiRateLimitIndex,
		func(rawObj k8client.Object) []string {
			ratelimitPolicy := rawObj.(*dpv1alpha1.RateLimitPolicy)
			var apis []string
			if ratelimitPolicy.Spec.TargetRef.Kind == constants.KindAPI {
				apis = append(apis,
					types.NamespacedName{
						Namespace: GetNamespace(
							(*gwapiv1b1.Namespace)(ratelimitPolicy.Spec.TargetRef.Namespace),
							ratelimitPolicy.Namespace),
						Name: string(ratelimitPolicy.Spec.TargetRef.Name),
					}.String())
			}
			logger.Info("index api policy")
			return apis
		})
	return err
}

// NamespacedName generates namespaced name for Kubernetes objects
func NamespacedName(obj client.Object) types.NamespacedName {
	return types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
}

// GetNamespace reads namespace with a default value
func GetNamespace(namespace *gwapiv1b1.Namespace, defaultNamespace string) string {
	if namespace != nil && *namespace != "" {
		return string(*namespace)
	}
	return defaultNamespace
}

// FilterByNamespaces takes a list of namespaces and returns a filter function
// which return true if the input object is in the given namespaces list,
// and returns false otherwise
func FilterByNamespaces(namespaces []string) func(object client.Object) bool {
	return func(object client.Object) bool {
		if namespaces == nil {
			return true
		}
		return stringutils.StringInSlice(object.GetNamespace(), namespaces)
	}
}

// GetOperatorPodNamespace returns the namesapce of the operator pod
func GetOperatorPodNamespace() string {
	return envutils.GetEnv(constants.OperatorPodNamespace,
		constants.OperatorPodNamespaceDefaultValue)
}

// SetupWithManager sets up the controller with the Manager.
func (ratelimitReconsiler *RateLimitPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dpv1alpha1.RateLimitPolicy{}).
		Complete(ratelimitReconsiler)
}
