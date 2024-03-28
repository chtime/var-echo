package main

import (
	"fmt"
	"net/http"
	"os"
)

func envHandler(w http.ResponseWriter, _ *http.Request) {
	chartName := os.Getenv("CHART_NAME")
	stage := os.Getenv("STAGE")
	namespace := os.Getenv("NAMESPACE")
	revision := os.Getenv("REVISION")
	message := "Hello from var-echo"
	_, _ = fmt.Fprintf(w, 
		"Chart Name: %s\nStage: %s\nNamespace: %s\nRevision: %s\nMessage: %s\n", 
		chartName, stage, namespace, revision, message
	)
}

func main() {
	http.HandleFunc("/env", envHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
