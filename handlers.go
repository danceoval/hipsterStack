package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "html"
    "io"
    "io/ioutil"

    "github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %q", html.EscapeString(r.URL.Path))
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(questions); err != nil{
        panic(err)
    }	
}

func SingleQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionId := vars["questionId"]
	fmt.Fprintln(w, "Question: ", questionId)
}

func QuestionCreate(w http.ResponseWriter, r *http.Request) {
    var question Question
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

    if err != nil {
        panic(err)
    }

    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &question); err !=nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := CreateQ(question)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}
