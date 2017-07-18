package main

import "fmt"

var currentId int

var questions Questions

func init() {
	CreateQ(Question{Q: "Who invented JavaScript?", Choices: [4]string{"Larry Page", "David Evans", "Brendan Eich", "Tim Heidecker",}, A: 2})
	CreateQ(Question{Q: "What's the shape of a Sierpinski Gasket?", Choices: [4]string{"Square", "Pentagram", "Snowflake", "Triangle",},A: 3})
	CreateQ(Question{Q: "Which of the following is not a GIRLS Girl?", Choices: [4]string{"Hannah", "Carrie", "Shoshana", "Marnie",},A: 1})
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