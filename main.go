package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello there, %s\n", r.URL.Path[1:])
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "[%s] Hi there, this is an response from \"%s\" with message: I love %s!\n", time.Now().String(), hostname, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
