//
// Copyright (c) 2022, WSO2 LLC. (http://www.wso2.com).
//
// WSO2 LLC. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

function getAPIList() returns string?|APIList|error {
    API[]|error? apis = db_getAPIsDAO();
    if apis is API[] {
        int count = apis.length();
        APIList apisList = {count: count, list: apis};
        return apisList;
    } else {
        return apis;
    }
}

function changeLifeCyleState(string action, string apiId, string organization) returns LifecycleState|error {
    string? prevLCState = check db_getCurrentLCStatus(apiId, organization);
    string|error? lcState = db_changeLCState(action, apiId, organization);
    if lcState is string {
            string? newvLCState = check db_getCurrentLCStatus(apiId, organization);
            string|error? lcEvent = db_AddLCEvent(apiId, prevLCState, newvLCState, organization);
            if lcEvent is string {
                json lcPayload = check getActionTransitionsFromState(action);
                LifecycleState lcCr = check lcPayload.cloneWithType(LifecycleState);
                return lcCr;
            } else {
                return error("error while adding LC event");
            }
    } else {
        return error("error while updating LC state");
    }
} 


function getLifeCyleState(string apiId, string organization) returns LifecycleState|error {
    string|error? currentLCState = db_getCurrentLCStatus(apiId, organization);
    if currentLCState is string {
        json lcPayload = check getTransitionsFromState(currentLCState);
        LifecycleState lcGet = check lcPayload.cloneWithType(LifecycleState);
        return lcGet;
    } else {
        return error("error while Getting LC state");
    }
}

function actionToLCState(any v) returns string {

    match v {
        "Demote to Created" => { return "CREATED"; }
        "Publish" => { return "PUBLISHED"; }
        "Block" => { return "BLOCKED"; }
        "Deprecate" => { return "DEPRECATED"; }
        "Retire" => { return "RETIRED"; }
        "Re-Publish" => { return "PUBLISHED"; }
        _ => { return "any"; }
    }
}

function getActionTransitionsFromState(string state) returns json|error {
    string actionState = actionToLCState(state);
    StatesList c =  check lifeCycleStateTransitions.cloneWithType(StatesList);
    foreach States x in c.States {
        if(actionState.equalsIgnoreCaseAscii(x.State)) {
            return x.toJson();
        }
    }
    
}

function getTransitionsFromState(string state) returns json|error {
    StatesList c =  check lifeCycleStateTransitions.cloneWithType(StatesList);
    foreach States x in c.States {
        if(state.equalsIgnoreCaseAscii(x.State)) {
            return x.toJson();
        }
    }
    
}

function getLcEventHistory(string apiId) returns LifecycleHistory|error {
    LifecycleHistoryItem[]|error? lcHistory = getLCEventHistory(apiId);
    if lcHistory is LifecycleHistoryItem[] {
        int count = lcHistory.length();
        LifecycleHistory eventList = {count: count, list: lcHistory};
        return eventList;
    } else {
        return error("Error while retriving LC events");
    }
}
