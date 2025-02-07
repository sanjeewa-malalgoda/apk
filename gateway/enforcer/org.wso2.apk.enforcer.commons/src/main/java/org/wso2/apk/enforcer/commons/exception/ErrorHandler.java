/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  WSO2 Inc. licenses this file to you under the Apache License,
 *  Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License.
 *  You may obtain a copy of the License at
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

package org.wso2.apk.enforcer.commons.exception;

/**
 * This error handler interface must use in all exceptions class, for example please see CommonGatewayException class.
 */
public interface ErrorHandler {
    /**
     * Get error code that defined in the enum.
     *
     * @return error code
     */
    long getErrorCode();

    /**
     * Get error message that defined in the enum.
     *
     * @return error message
     */
    String getErrorMessage();

    /**
     * Get error description that defined in the enum.
     *
     * @return error description.
     */
    String getErrorDescription();

    /**
     * Get Http status code that defined in the enum.
     *
     * @return error code.
     */
    int getHttpStatusCode();

    /**
     * Return true if stack trace to print.
     *
     * @return
     */
    boolean printStackTrace();
}
