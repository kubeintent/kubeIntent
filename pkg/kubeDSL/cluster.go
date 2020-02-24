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

import (
	"github.com/adibrastegarnia/kubeDSL/pkg/kube"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NewCluster creates a new cluster builder
func NewCluster(namespace string) ClusterBuilder {
	kubeApi, err := kube.GetAPI(namespace)
	if err != nil {
		panic(err)
	}
	return &Cluster{
		kube: kubeApi,
	}
}

type ClusterBuilder interface {
	SetPods(...Pod) ClusterBuilder
	SetDeployments(...Deployment) ClusterBuilder
	Build() Cluster
}

// Cluster k8s cluster type
type Cluster struct {
	*client
	kube        kube.API
	Pods        []Pod
	Deployments []Deployment
}

// SetDeployments sets cluster deployments
func (c *Cluster) SetDeployments(deployments ...Deployment) ClusterBuilder {
	c.Deployments = deployments
	return c
}

// SetPods set cluster pods
func (c *Cluster) SetPods(pods ...Pod) ClusterBuilder {
	c.Pods = pods
	return c
}

// Build builds a k8s cluster
func (c *Cluster) Build() Cluster {
	client := &client{
		namespace:        c.kube.Namespace(),
		config:           c.kube.Config(),
		kubeClient:       kubernetes.NewForConfigOrDie(c.kube.Config()),
		extensionsClient: apiextension.NewForConfigOrDie(c.kube.Config()),
	}
	return Cluster{
		client:      client,
		Pods:        c.Pods,
		Deployments: c.Deployments,
	}
}

func (c *Cluster) createDeployments() error {
	for _, deployment := range c.Deployments {
		var containers []corev1.Container
		for _, container := range deployment.pod.containers {
			var ports []corev1.ContainerPort
			for _, port := range container.ports {
				ports = append(ports, corev1.ContainerPort{
					Name:          port.name,
					ContainerPort: port.containerPort,
					HostIP:        port.hostIP,
					HostPort:      port.hostPort,
					Protocol:      corev1.Protocol(port.protocol),
				})
			}
			containers = append(containers, corev1.Container{
				Name:            container.name,
				Image:           container.image,
				Args:            container.args,
				Command:         container.command,
				ImagePullPolicy: corev1.PullPolicy(container.pullPolicy),
				Ports:           ports,
			})
		}
		kubeDeployment := appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      deployment.name,
				Namespace: c.namespace,
				Labels:    deployment.labels,
			},

			Spec: appsv1.DeploymentSpec{
				Replicas: &deployment.replicas,
				Selector: &metav1.LabelSelector{
					MatchLabels: deployment.labels,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: deployment.labels,
					},
					Spec: corev1.PodSpec{
						Containers: containers,
					},
				},
			},
		}

		_, err := c.client.kubeClient.AppsV1().Deployments(c.namespace).Create(&kubeDeployment)
		if err != nil {
			return err
		}

	}

	return nil
}

// createPods creates cluster pods
func (c *Cluster) createPods() error {
	for _, pod := range c.Pods {
		var containers []corev1.Container
		for _, container := range pod.containers {
			var ports []corev1.ContainerPort
			for _, port := range container.ports {
				ports = append(ports, corev1.ContainerPort{
					Name:          port.name,
					ContainerPort: port.containerPort,
					HostIP:        port.hostIP,
					HostPort:      port.hostPort,
					Protocol:      corev1.Protocol(port.protocol),
				})
			}
			containers = append(containers, corev1.Container{
				Name:            container.name,
				Image:           container.image,
				Args:            container.args,
				Command:         container.command,
				ImagePullPolicy: corev1.PullPolicy(container.pullPolicy),
				Ports:           ports,
			})
		}
		kubePod := corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      pod.name,
				Namespace: c.namespace,
			},
			Spec: corev1.PodSpec{
				Containers: containers,
			},
		}
		_, err := c.client.kubeClient.CoreV1().Pods(c.namespace).Create(&kubePod)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cluster) CreateCluster() error {
	// create a set of pods
	if len(c.Pods) != 0 {
		err := c.createPods()
		if err != nil {
			return err
		}
	}
	// create a set of deployments
	if len(c.Deployments) != 0 {
		err := c.createDeployments()
		if err != nil {
			return err
		}
	}

	// create a set of deployments
	return nil
}
