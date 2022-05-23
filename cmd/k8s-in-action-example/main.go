package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func logRequest(req *http.Request) {
	log.Printf("Req: %s %s", req.Method, req.URL)
}

func main() {

	catchallHandler := func(w http.ResponseWriter, req *http.Request) {
		logRequest(req)
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 Not Found\n")
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		logRequest(req)
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("getting hostname failed: %v", err)
			hostname = "unknown host"
		}
		res := fmt.Sprintf("Hello from %s\n", hostname)
		io.WriteString(w, res)
	}

	http.HandleFunc("/", catchallHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening for requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
