package kube

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	// v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getClientset() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to kubeconfig
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("can't load kubeconfig: %w", err)
		}
	}
	return kubernetes.NewForConfig(config)
}

func ListAllPodLabels() error {
	clientset, err := getClientset()
	if err != nil {
		return err
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod: %s/%s\n", pod.Namespace, pod.Name)
		printLabels(pod.Labels)
		fmt.Println("----")
	}

	return nil
}

func printLabels(labels map[string]string) {
	if len(labels) == 0 {
		fmt.Println("No labels")
		return
	}
	for k, v := range labels {
		fmt.Printf(" %s = %s\n", k, v)
	}
}
