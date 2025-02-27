// Code generated by protoc. DO NOT EDIT.
// Copyright (c) 2021, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package server

import (
	"context"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	envoy_delta "github.com/envoyproxy/go-control-plane/pkg/server/delta/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/rest/v3"
	envoy_sotw "github.com/envoyproxy/go-control-plane/pkg/server/sotw/v3"
	streamv3 "github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
	xdsv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/api"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/apkmgt"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/config"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/keymgt"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/subscription"
	"github.com/wso2/apk/adapter/pkg/discovery/protocol/resource/v3"
	"github.com/wso2/apk/adapter/pkg/discovery/protocol/server/sotw/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is a collection of handlers for streaming discovery requests.
type Server interface {
	config.ConfigDiscoveryServiceServer
	api.ApiDiscoveryServiceServer
	subscription.SubscriptionDiscoveryServiceServer
	subscription.ApplicationDiscoveryServiceServer
	subscription.ApiListDiscoveryServiceServer
	subscription.ApplicationPolicyDiscoveryServiceServer
	subscription.JWTIssuerDiscoveryServiceServer
	subscription.SubscriptionPolicyDiscoveryServiceServer
	subscription.ApplicationKeyMappingDiscoveryServiceServer
	keymgt.KMDiscoveryServiceServer
	keymgt.RevokedTokenDiscoveryServiceServer
	apkmgt.APKMgtDiscoveryServiceServer

	rest.Server
	envoy_sotw.Server
	envoy_delta.Server
}

// NewServer creates handlers from a config watcher and callbacks.
func NewServer(ctx context.Context, config cache.Cache, callbacks xdsv3.Callbacks) Server {
	// Delta server is not used. Hence the envoy default implementation is used.
	return NewServerAdvanced(rest.NewServer(config, callbacks), sotw.NewServer(ctx, config, callbacks), envoy_delta.NewServer(ctx, config, callbacks))
}

// NewServerAdvanced creates new server object
func NewServerAdvanced(restServer rest.Server, sotwServer envoy_sotw.Server, deltaServer envoy_delta.Server) Server {
	return &server{rest: restServer, sotw: sotwServer, delta: deltaServer}
}

type server struct {
	config.UnimplementedConfigDiscoveryServiceServer
	api.UnimplementedApiDiscoveryServiceServer
	subscription.UnimplementedSubscriptionDiscoveryServiceServer
	subscription.UnimplementedApplicationDiscoveryServiceServer
	subscription.UnimplementedJWTIssuerDiscoveryServiceServer
	subscription.UnimplementedApiListDiscoveryServiceServer
	subscription.UnimplementedApplicationPolicyDiscoveryServiceServer
	subscription.UnimplementedSubscriptionPolicyDiscoveryServiceServer
	subscription.UnimplementedApplicationKeyMappingDiscoveryServiceServer
	keymgt.UnimplementedKMDiscoveryServiceServer
	keymgt.UnimplementedRevokedTokenDiscoveryServiceServer
	apkmgt.UnimplementedAPKMgtDiscoveryServiceServer
	rest  rest.Server
	sotw  envoy_sotw.Server
	delta envoy_delta.Server
}

func (s *server) StreamHandler(stream streamv3.Stream, typeURL string) error {
	return s.sotw.StreamHandler(stream, typeURL)
}

func (s *server) StreamConfigs(stream config.ConfigDiscoveryService_StreamConfigsServer) error {
	return s.StreamHandler(stream, resource.ConfigType)
}

func (s *server) StreamApis(stream api.ApiDiscoveryService_StreamApisServer) error {
	return s.StreamHandler(stream, resource.APIType)
}

func (s *server) StreamSubscriptions(stream subscription.SubscriptionDiscoveryService_StreamSubscriptionsServer) error {
	return s.StreamHandler(stream, resource.SubscriptionListType)
}

func (s *server) StreamApiList(stream subscription.ApiListDiscoveryService_StreamApiListServer) error {
	return s.StreamHandler(stream, resource.APIListType)
}

func (s *server) StreamApplications(stream subscription.ApplicationDiscoveryService_StreamApplicationsServer) error {
	return s.StreamHandler(stream, resource.ApplicationListType)
}

func (s *server) StreamApplicationPolicies(stream subscription.ApplicationPolicyDiscoveryService_StreamApplicationPoliciesServer) error {
	return s.StreamHandler(stream, resource.ApplicationPolicyListType)
}

func (s *server) StreamSubscriptionPolicies(stream subscription.SubscriptionPolicyDiscoveryService_StreamSubscriptionPoliciesServer) error {
	return s.StreamHandler(stream, resource.SubscriptionPolicyListType)
}

func (s *server) StreamApplicationKeyMappings(stream subscription.ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsServer) error {
	return s.StreamHandler(stream, resource.ApplicationKeyMappingListType)
}

func (s *server) StreamKeyManagers(stream keymgt.KMDiscoveryService_StreamKeyManagersServer) error {
	return s.StreamHandler(stream, resource.KeyManagerType)
}

func (s *server) StreamTokens(stream keymgt.RevokedTokenDiscoveryService_StreamTokensServer) error {
	return s.StreamHandler(stream, resource.RevokedTokensType)
}

func (s *server) StreamAPKMgtApplications(stream apkmgt.APKMgtDiscoveryService_StreamAPKMgtApplicationsServer) error {
	return s.StreamHandler(stream, resource.APKMgtApplicationType)
}
func (s *server)StreamJWTIssuers(stream subscription.JWTIssuerDiscoveryService_StreamJWTIssuersServer) error {
	return s.StreamHandler(stream, resource.JWTIssuerListType)
}

// Fetch is the universal fetch method.
func (s *server) Fetch(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	return s.rest.Fetch(ctx, req)
}

func (s *server) FetchConfigs(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.ConfigType
	return s.Fetch(ctx, req)
}

func (s *server) FetchApis(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Unauthenticated, "empty request")
	}
	req.TypeUrl = resource.APIType
	return s.Fetch(ctx, req)
}

func (s *server) FetchTokens(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.RevokedTokensType
	return s.Fetch(ctx, req)
}

func (s *server) FetchThrottleData(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.ThrottleDataType
	return s.Fetch(ctx, req)
}

func (s *server) FetchAPKMgtApplications(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.APKMgtApplicationType
	return s.Fetch(ctx, req)
}

func (s *server) DeltaStreamHandler(stream streamv3.DeltaStream, typeURL string) error {
	return s.delta.DeltaStreamHandler(stream, typeURL)
}
