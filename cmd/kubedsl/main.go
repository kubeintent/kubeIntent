package main

import (
	"github.com/adibrastegarnia/kubeDSL/pkg/kube"
	"github.com/adibrastegarnia/kubeDSL/pkg/kubeDSL"
)

func main() {
	api, err := kube.GetAPI("default")
	if err != nil {
		panic(err)
	}

	pod := kubeDSL.NewPod()
	pod.SetName("hello-world")

	cluster := kubeDSL.NewCluster(api)
	c := cluster.SetPods(pod.Build()).Build()
	err = c.CreatePods()
	if err != nil {
		panic(err)
	}

}
