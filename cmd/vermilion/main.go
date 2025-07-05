package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dtkachenko/vermilion/internal/handler"
	"github.com/dtkachenko/vermilion/internal/kube"
	"github.com/dtkachenko/vermilion/internal/storage"
	"github.com/dtkachenko/vermilion/internal/storage/memory"
)

func main() {
	fmt.Println("Service Start")

	go func() {
		if err := kube.ListAllPodLabels(); err != nil {
			log.Fatalf("Failed to list pods: %v", err)
		}
	}()

	pods, err := kube.GetAllPodsLabels()
	if err != nil {
		log.Fatalf("Can't get pods %v", err)
	}
	s := getStore()

	for _, pod := range pods.Items {
		s.Save(storage.PodInfo{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Labels:    pod.Labels,
		})
	}

	storedPods, err := s.GetAll()
	if err != nil {
		fmt.Printf("%v", err)
	}

	jsonData, err := json.MarshalIndent(storedPods, "", "  ")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf(string(jsonData)) 

	http.HandleFunc("/", handler.HelloHandler)

	log.Println("Starting serivce on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getStore() storage.Store {
	switch "memory" {
	case "memory":
		return memory.New()
	default:
		log.Fatal("Wrong storage type in VERMILION_STORAGE_BACKEND")
		return nil
	}
}
