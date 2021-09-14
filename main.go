package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func indexRoute(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Welcome to my API")
}

func getTasks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(tasks)
}

func getTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(rw, "Invalid ID")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(task)
		}
	}
}

func createTask(rw http.ResponseWriter, r *http.Request) {
	var task Task
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(rw, "Insert a valid task")
	}
	json.Unmarshal(requestBody, &task)
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(task)
}

func updateTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(rw, "Invalid ID")
		return
	}

	var updatedTask Task
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(rw, "Enter valid data")
		return
	}
	json.Unmarshal(requestBody, &updatedTask)

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(rw, "The task with ID %v has been succcessfully updated", taskID)
		}
	}
}

func deleteTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(rw, "Invalid ID")
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Fprintf(rw, "The task with %v has been removed successfully", taskID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
