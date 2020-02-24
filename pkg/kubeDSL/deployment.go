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

// NewDeployment creates a new deployment
func NewDeployment() DeploymentBuilder {
	return &Deployment{}
}

type DeploymentBuilder interface {
	SetName(string) DeploymentBuilder
	SetPods(Pod) DeploymentBuilder
	SetReplicas(int32) DeploymentBuilder
	Build() Deployment
}

// Deployment deployment abstraction
type Deployment struct {
	name     string
	pod      Pod
	replicas int32
}

// Replicas returns the number of deployment replicas
func (d *Deployment) Replicas() int32 {
	return d.replicas
}

// SetReplicas sets deployment replicas
func (d *Deployment) SetReplicas(replicas int32) DeploymentBuilder {
	d.replicas = replicas
	return d
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

// SetPods sets deployment pod
func (d *Deployment) SetPods(pod Pod) DeploymentBuilder {
	d.pod = pod
	return d
}

// Pods returns deployment pod
func (d *Deployment) Pod() Pod {
	return d.pod
}

// Build builds a deployment
func (d *Deployment) Build() Deployment {
	return Deployment{
		name:     d.name,
		pod:      d.pod,
		replicas: d.replicas,
	}
}
