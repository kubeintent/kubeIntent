// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubeDSL

type DeploymentBuilder interface {
	SetName(string) DeploymentBuilder
	SetPods(...Pod) DeploymentBuilder
	Build() Deployment
}

// Deployment deployment abstraction
type Deployment struct {
	name string
	pods []Pod
}

// SetName sets a deployment name
func (d *Deployment) SetName(name string) DeploymentBuilder {
	d.name = name
	return d
}

// Name returns the name of a deployment
func (d *Deployment) Name() string {
	return d.name
}

// SetPods sets deployment pods
func (d *Deployment) SetPods(pods ...Pod) DeploymentBuilder {
	d.pods = pods
	return d
}

// Pods returns deployment pods
func (d *Deployment) Pods() []Pod {
	return d.pods
}

// Build builds a deployment
func (d *Deployment) Build() Deployment {
	return Deployment{
		name: d.name,
		pods: d.pods,
	}
}
