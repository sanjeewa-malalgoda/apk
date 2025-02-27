/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package org.wso2.apk.enforcer.commons.analytics.publishers.dto;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import org.wso2.apk.enforcer.commons.analytics.publishers.dto.Error;
import java.util.Map;

/**
 * analytics event data.
 */
public class Event {
    @JsonUnwrapped
    private API api;
    @JsonUnwrapped
    private Operation operation;
    @JsonUnwrapped
    private Target target;
    @JsonUnwrapped
    private Application application;
    @JsonUnwrapped
    private Latencies latencies;
    @JsonUnwrapped
    private MetaInfo metaInfo;
    @JsonUnwrapped
    private Error error;
    private int proxyResponseCode;
    private String requestTimestamp;
    private String userAgentHeader;

    private String userName;
    private String userIp;

    private String errorType;

    private Map<String, Object> properties;

    public API getApi() {
        return api;
    }

    public void setApi(API api) {
        this.api = api;
    }

    public Operation getOperation() {
        return operation;
    }

    public void setOperation(Operation operation) {
        this.operation = operation;
    }

    public Target getTarget() {
        return target;
    }

    public void setTarget(Target target) {
        this.target = target;
    }

    public Application getApplication() {
        return application;
    }

    public void setApplication(Application application) {
        this.application = application;
    }

    public Latencies getLatencies() {
        return latencies;
    }

    public void setLatencies(Latencies latencies) {
        this.latencies = latencies;
    }

    public int getProxyResponseCode() {
        return proxyResponseCode;
    }

    public void setProxyResponseCode(int proxyResponseCode) {
        this.proxyResponseCode = proxyResponseCode;
    }

    public String getRequestTimestamp() {
        return requestTimestamp;
    }

    public void setRequestTimestamp(String requestTimestamp) {
        this.requestTimestamp = requestTimestamp;
    }

    public Error getError() {
        return error;
    }

    public void setError(Error error) {
        this.error = error;
    }

    public String getErrorType() {
        return errorType;
    }

    public void setErrorType(String errorType) {
        this.errorType = errorType;
    }

    public MetaInfo getMetaInfo() {
        return metaInfo;
    }

    public void setMetaInfo(MetaInfo metaInfo) {
        this.metaInfo = metaInfo;
    }

    public String getUserAgentHeader() {
        return userAgentHeader;
    }

    public void setUserAgentHeader(String userAgentHeader) {
        this.userAgentHeader = userAgentHeader;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getUserIp() {
        return userIp;
    }

    public void setUserIp(String userIp) {
        this.userIp = userIp;
    }

    public Map<String, Object> getProperties() {
        return properties;
    }

    public void setProperties(Map<String, Object> properties) {
        this.properties = properties;
    }
}
