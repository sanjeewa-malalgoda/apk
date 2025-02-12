/*
 * Copyright (c) 2018, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.wso2.apk.enforcer.listener.events;

import org.wso2.apk.enforcer.constants.APIConstants;

import java.util.Objects;

/**
 * An Event Object which can holds the data related to API Policy which are required
 * for the validation purpose in a gateway.
 */
public class APIPolicyEvent extends PolicyEvent {
    private int policyId;
    private String policyName;
    private String quotaType;

    public APIPolicyEvent(String eventId, long timestamp, String type, int tenantId, String tenantDomain, int policyId,
                          String policyName, String quotaType) {
        this.eventId = eventId;
        this.timeStamp = timestamp;
        this.type = type;
        this.tenantId = tenantId;
        this.policyId = policyId;
        this.policyName = policyName;
        this.quotaType = quotaType;
        this.tenantDomain = tenantDomain;
        this.policyType = APIConstants.PolicyType.API;
    }

    @Override
    public String toString() {
        return "APIPolicyEvent{" +
                "policyId=" + policyId +
                ", policyName='" + policyName + '\'' +
                ", quotaType='" + quotaType + '\'' +
                ", eventId='" + eventId + '\'' +
                ", timeStamp=" + timeStamp +
                ", type='" + type + '\'' +
                ", tenantId=" + tenantId + '\'' +
                ", tenantDomain=" + tenantDomain +
                '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (!(o instanceof APIPolicyEvent)) {
            return false;
        }
        APIPolicyEvent that = (APIPolicyEvent) o;
        return getPolicyId() == that.getPolicyId() &&
                getPolicyName().equals(that.getPolicyName()) &&
                getQuotaType().equals(that.getQuotaType());
    }

    @Override
    public int hashCode() {
        return Objects.hash(getPolicyId(), getPolicyName(), getQuotaType());
    }

    public int getPolicyId() {
        return policyId;
    }

    public void setPolicyId(int policyId) {
        this.policyId = policyId;
    }

    public String getPolicyName() {
        return policyName;
    }

    public void setPolicyName(String policyName) {
        this.policyName = policyName;
    }

    public String getQuotaType() {
        return quotaType;
    }

    public void setQuotaType(String quotaType) {
        this.quotaType = quotaType;
    }    
}
