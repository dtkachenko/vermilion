package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dtkachenko/vermilion/internal/handler"
	"github.com/dtkachenko/vermilion/internal/kube"
)

func main() {
	fmt.Println("Service Start")

	go func() {
		if err := kube.ListAllPodLabels(); err != nil {
			log.Fatalf("Failed to list pods: %v", err)
		}
	}()

	http.HandleFunc("/", handler.HelloHandler)
	
	log.Println("Starting serivce on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
