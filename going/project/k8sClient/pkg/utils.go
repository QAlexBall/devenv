package pkg

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfig() (string, error) {
	return "/workspace/.kube/config", nil
}

func GetKubernetesClient() (*kubernetes.Clientset, error) {
	kubeconfig, err := getKubeConfig()
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
