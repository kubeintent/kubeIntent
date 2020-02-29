package main

import (
	"github.com/adibrastegarnia/kubeDSL/pkg/intent"
)

func main() {

	lables := make(map[string]string)
	lables["app"] = "hello-world-deployment"

	container := intent.NewContainer()
	container.SetName("hello-world-8")
	container.SetImage("callicoder/go-hello-world:1.0.0")
	container.SetPullPolicy(intent.Always.String())
	containerInst := container.Build()

	pod := intent.NewPod()
	pod.SetName("hello-world-8")
	pod.SetContainers(containerInst)
	podInst := pod.Build()

	deployment := intent.NewDeployment()
	deployment.SetName("hello-world-deployment").SetLabels(lables)
	deploymentInst := deployment.SetReplicas(2).SetPods(podInst).Build()

	cluster := intent.NewCluster("kube-dsl-2")
	clusterInst := cluster.SetDeployments(deploymentInst).Build()
	err := clusterInst.CreateCluster()
	if err != nil {
		panic(err)
	}

}
