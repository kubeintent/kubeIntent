package main

import (
	"github.com/adibrastegarnia/kubeDSL/pkg/kubeDSL"
)

func main() {

	lables := make(map[string]string)
	lables["app"] = "hello-world-deployment"

	container := kubeDSL.NewContainer()
	container.SetName("hello-world-8")
	container.SetImage("callicoder/go-hello-world:1.0.0")
	container.SetPullPolicy(kubeDSL.Always.String())
	containerInst := container.Build()

	pod := kubeDSL.NewPod()
	pod.SetName("hello-world-8")
	pod.SetContainers(containerInst)
	podInst := pod.Build()

	deployment := kubeDSL.NewDeployment()
	deployment.SetName("hello-world-deployment").SetLabels(lables)
	deploymentInst := deployment.SetReplicas(2).SetPods(podInst).Build()

	cluster := kubeDSL.NewCluster("kube-dsl")
	clusterInst := cluster.SetDeployments(deploymentInst).Build()
	err := clusterInst.CreateCluster()
	if err != nil {
		panic(err)
	}

}
