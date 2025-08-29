package main

import (
    "log"
    "net/http"
)

func main() {
    InitDB()

    http.Handle("/", http.FileServer(http.Dir("../frontend")))
    http.HandleFunc("/api/register", RegisterHandler)
    http.HandleFunc("/api/login", LoginHandler)

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
