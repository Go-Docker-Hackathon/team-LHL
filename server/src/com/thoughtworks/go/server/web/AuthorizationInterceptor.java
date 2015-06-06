/*************************GO-LICENSE-START*********************************
 * Copyright 2014 ThoughtWorks, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *************************GO-LICENSE-END***********************************/

package com.thoughtworks.go.server.web;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import static javax.servlet.http.HttpServletResponse.SC_UNAUTHORIZED;

import com.thoughtworks.go.config.CaseInsensitiveString;
import com.thoughtworks.go.server.service.SecurityService;
import com.thoughtworks.go.server.util.UserHelper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

@Controller
public class AuthorizationInterceptor implements HandlerInterceptor {
    private SecurityService securityService;

    @Autowired
    public AuthorizationInterceptor(SecurityService securityService) {
        this.securityService = securityService;
    }

    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
            throws Exception {
        if (UserHelper.isAgent()) {
            return true;
        }

        String pipelineName = request.getParameter("pipelineName");
        if (pipelineName != null) {
            String userName = CaseInsensitiveString.str(UserHelper.getUserName().getUsername());
            if (request.getMethod().equalsIgnoreCase("get")) {
                if (!securityService.hasViewPermissionForPipeline(userName, pipelineName)) {
                    response.sendError(SC_UNAUTHORIZED);
                    return false;
                }
            } else if (request.getMethod().equalsIgnoreCase("post") || request.getMethod().equalsIgnoreCase("put")) {
                if (isEditingConfigurationRequest(request)) {
                    return true;
                }
                
                String stageName = request.getParameter("stageName");
                if (stageName != null) {
                    if (!securityService.hasOperatePermissionForStage(pipelineName, stageName, userName)) {
                        response.sendError(SC_UNAUTHORIZED);
                        return false;
                    }
                } else {
                    if (isForcePipelineRequest(request)) {
                        if (!securityService.hasOperatePermissionForFirstStage(pipelineName, userName)) {
                            response.sendError(SC_UNAUTHORIZED);
                            return false;
                        }
                    } else if (!securityService.hasOperatePermissionForPipeline(new CaseInsensitiveString(userName), pipelineName)) {
                        response.sendError(SC_UNAUTHORIZED);
                        return false;
                    }
                }
            }
        }
        return true;
    }

    private boolean isEditingConfigurationRequest(HttpServletRequest request) {
        return request.getRequestURI().indexOf("/admin/restful/configuration") != -1;
    }

    private boolean isForcePipelineRequest(HttpServletRequest request) {
        return request.getRequestURI().endsWith("/force");
    }

    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler,
                           ModelAndView modelAndView) throws Exception {
    }

    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex)
            throws Exception {
    }
}
