package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/questions", Questions)
    router.HandleFunc("/questions/{questionId}", SingleQuestion)

    fmt.Printf("*** STARTING UP ON PORT 8080! *** \n")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Questions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Question Index")
}

func SingleQuestion(w http.ResponseWriter, r *http.Request) {
	vars :=mux.Vars(r)
	questionId := vars["questionId"]
	fmt.Fprintln(w, "Question: ", questionId)
}