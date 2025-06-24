package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "✅ Hello from Golang service on /service1!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting Go server on port 8001...")
    log.Fatal(http.ListenAndServe(":8001", nil))
}
