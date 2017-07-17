package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    router := NewRouter()

    fmt.Printf("*** STARTING UP ON PORT 8080! *** \n")
    log.Fatal(http.ListenAndServe(":8080", router))
}
