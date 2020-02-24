package main

import (
	"github.com/adibrastegarnia/kubeDSL/pkg/kubeDSL"
)

func main() {

	container := kubeDSL.NewContainer()
	container.SetName("hello-world-6")
	container.SetImage("callicoder/go-hello-world:1.0.0")
	container.SetPullPolicy(kubeDSL.Always.String())
	containerInst := container.Build()

	pod := kubeDSL.NewPod()
	pod.SetName("hello-world-6")
	pod.SetContainers(containerInst)
	podInst := pod.Build()

	deployment := kubeDSL.NewDeployment()
	deployment.SetName("hello-world-deployment")
	deploymentInst := deployment.SetReplicas(2).SetPods(podInst).Build()

	cluster := kubeDSL.NewCluster("default")
	clusterInst := cluster.SetPods(podInst).SetDeployments(deploymentInst).Build()
	err := clusterInst.CreateCluster()
	if err != nil {
		panic(err)
	}

}
