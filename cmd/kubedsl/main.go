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

	/*container2 := kubeDSL.NewContainer()
	container2.SetName("hello-world-5-2")
	container2.SetImage("callicoder/go-hello-world:1.0.0")
	container2.SetPullPolicy(kubeDSL.Always.String())
	containerInst2 := container2.Build()*/

	pod := kubeDSL.NewPod()
	pod.SetName("hello-world-6")
	pod.SetContainers(containerInst)
	podInst := pod.Build()

	cluster := kubeDSL.NewCluster("default")
	clusterInst := cluster.SetPods(podInst).Build()
	err := clusterInst.CreateCluster()
	if err != nil {
		panic(err)
	}

}
