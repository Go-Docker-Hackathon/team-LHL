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

package com.thoughtworks.go.config.preprocessor;

import com.thoughtworks.go.config.CruiseConfig;
import com.thoughtworks.go.config.GoConfigPreprocessor;
import com.thoughtworks.go.config.ParamsConfig;

/**
 * @understands Interpolation of config parameters
 */
public class ConfigParamPreprocessor implements GoConfigPreprocessor {
    private ClassAttributeCache.FieldCache fieldCache;
    private ClassAttributeCache.AnnotationPresentCache annotationPresentCache;


    public ConfigParamPreprocessor() {
        fieldCache = new ClassAttributeCache.FieldCache();
        annotationPresentCache = new ClassAttributeCache.AnnotationPresentCache();
    }

    public void process(CruiseConfig cruiseConfig) {
        ParamResolver resolver = new ParamResolver(new ParamSubstitutionHandlerFactory(new ParamsConfig()), fieldCache);
        resolver.resolve(cruiseConfig);
    }

}
