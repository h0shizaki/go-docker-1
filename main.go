package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Status struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	currentStatus := Status{
		Status:      "Available",
		Environment: os.Getenv("Environment"),
	}
	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func main() {

	if os.Getenv("Environment") == "" {
		os.Setenv("Environment", "Develop")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("Hello from go lang"))
	})

	http.HandleFunc("/status", checkStatus)

	fmt.Println("SERVER IS RUNNING ON PORT 8080")
	http.ListenAndServe(":8080", nil)
}
