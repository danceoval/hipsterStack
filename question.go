package main 

type Question struct {
	Id int `json:"id"`
	Q string `json:"question"`
	A string `json:"answer"`
	Answered bool `json:"answered"`
}

type Questions []Question