package main

import (
	"fmt"
	"net/http"
	"os"
)

func envHandler(w http.ResponseWriter, _ *http.Request) {
	envVar := os.Getenv("LE_ENV_VAR")
	_, _ = fmt.Fprintf(w, "Value of LE_ENV_VAR is: %s", envVar)
}

func main() {
	http.HandleFunc("/env", envHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
