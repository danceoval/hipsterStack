package main 

type Question struct {
	Id int `json:"id"`
	Q string `json:"question"`
	Choices [4]string `json:"choices"`
	A int `json:"answer"`
	Answered bool `json:"answered"`
}

type Questions []Question