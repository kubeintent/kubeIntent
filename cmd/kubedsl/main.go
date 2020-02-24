package main

import (
	"github.com/adibrastegarnia/kubeDSL/pkg/kubeDSL"
)

func main() {

	container := kubeDSL.NewContainer()
	container.SetName("hello-world-3")
	container.SetImage("callicoder/go-hello-world:1.0.0")
	container.SetPullPolicy("Always")
	containerInst := container.Build()

	pod := kubeDSL.NewPod()
	pod.SetName("hello-world-3")
	pod.SetContainers(containerInst)
	podInst := pod.Build()

	cluster := kubeDSL.NewCluster("default")
	clusterInst := cluster.SetPods(podInst).Build()
	err := clusterInst.CreateCluster()
	if err != nil {
		panic(err)
	}

}
