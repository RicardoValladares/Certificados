package main

import (
    "fmt"
    "github.com/kabukky/httpscerts"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola!")
}

func main() {
    err := httpscerts.Check("cert.pem", "key.pem")
    if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", "demo.com")
        if err != nil {
            log.Fatal("Error: No se logro crear https certificado.")
        }
    }
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}
