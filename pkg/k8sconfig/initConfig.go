package k8sconfig

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func K8sConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", "./resource/config")
	if err != nil {
		log.Fatal(err)
	}
	config.Insecure = true
	return config
}
