package main

import (
    "log" 
    "net/http"
)

func main() {

    fs := http.FileServer(http.Dir("public"))
    http.Handle("/", fs)

    log.Println("Listening to Port3000...")

    http.ListenAndServe(":3000", nil)

}