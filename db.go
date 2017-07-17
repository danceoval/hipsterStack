package main

import "fmt"

var currentId int

var questions Questions

func init() {
	CreateQ(Question{Q: "What's black, white and read all over?", A: "A Newspaper"})
	CreateQ(Question{Q: "What's the difference between a piano and a fish?", A: "You can't tuna fish"})
}

func FindQ(id int) Question {

	for _, q := range questions {
		if q.Id == id {
			return q
		}
	}

	return Question{}
}

func CreateQ(q Question) Question {
	currentId += 1
	q.Id = currentId
	questions = append(questions, q)
	return q
}

func DestoryQ(id int) error {
	for i, q := range questions {
		if q.Id == id {
			questions = append(questions[:i], questions[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find question qith id of %d", id)
}