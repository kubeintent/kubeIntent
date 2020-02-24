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
	corev1 "k8s.io/api/core/v1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NewCluster creates a new cluster builder
func NewCluster(kube kube.API) ClusterBuilder {
	return &Cluster{
		kube: kube,
	}
}

type ClusterBuilder interface {
	SetPods(...Pod) ClusterBuilder
	Build() Cluster
}

// Cluster k8s cluster type
type Cluster struct {
	*client
	kube kube.API
	Pods []Pod
}

func (c *Cluster) SetPods(pods ...Pod) ClusterBuilder {
	c.Pods = pods
	return c
}

func (c *Cluster) Build() Cluster {
	client := &client{
		namespace:        c.kube.Namespace(),
		config:           c.kube.Config(),
		kubeClient:       kubernetes.NewForConfigOrDie(c.kube.Config()),
		extensionsClient: apiextension.NewForConfigOrDie(c.kube.Config()),
	}
	return Cluster{
		client: client,
		Pods:   c.Pods,
	}
}

func (c *Cluster) CreatePods() error {
	for _, pod := range c.Pods {
		kubePod := corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      pod.name,
				Namespace: c.namespace,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            pod.name,
						Image:           "callicoder/go-hello-world:1.0.0",
						ImagePullPolicy: "Always",
					},
				},
			},
		}
		_, err := c.client.kubeClient.CoreV1().Pods(c.namespace).Create(&kubePod)
		if err != nil {
			return err
		}
	}
	return nil
}
