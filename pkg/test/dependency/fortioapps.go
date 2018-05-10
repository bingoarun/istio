//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package dependency

import (
	"istio.io/istio/pkg/test/environment"
	"istio.io/istio/pkg/test/internal"
)

// FortioApps indicates a dependency on FortioApps.
var FortioApps Dependency = &fortioapps{}

type fortioapps struct {
}

var _ Dependency = &fortioapps{}
var _ internal.Stateful = &fortioapps{}

func (f *fortioapps) String() string {
	return ""
}

func (f *fortioapps) Initialize(env environment.Interface) (interface{}, error) {
	return nil, nil
}

func (f *fortioapps) Reset(env environment.Interface, state interface{}) error {
	return nil
}

func (f *fortioapps) Cleanup(env environment.Interface, state interface{}) {

}
