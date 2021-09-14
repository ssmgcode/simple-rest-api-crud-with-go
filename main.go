package main

import "github.com/gorilla/mux"

type Task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type AllTasks []Task

var tasks AllTasks = AllTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some content",
	},
}

func main() {
	mux.NewRouter()
}
