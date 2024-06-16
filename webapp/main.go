package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	println("Starting webserver")
	port := "0.0.0.0:8000"
	originServerHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[origin server] received request at: %s\n", time.Now())
		_, _ = fmt.Fprint(rw, "origin server response")
	})
	log.Println("Listening on Port ", port)

	log.Fatal(http.ListenAndServe(port, originServerHandler))
}
