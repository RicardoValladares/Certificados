package main

import (
    "fmt"
    "github.com/kabukky/httpscerts"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

func main() {
    // Check if the cert files are available.
    err := httpscerts.Check("cert.pem", "key.pem")
    // If they are not available, generate new ones.
    if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", "demo.com")
        if err != nil {
            log.Fatal("Error: Couldn't create https certs.")
        }
    }
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}
