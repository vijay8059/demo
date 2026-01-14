package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("could not get hostname: %v", err)
		hostname = "unknown"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request %s %s from %s served by pod/container %s",
			r.Method, r.URL.Path, r.RemoteAddr, hostname)

		fmt.Fprintf(w, "Hello, World! from %s\n", hostname)
	})

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
