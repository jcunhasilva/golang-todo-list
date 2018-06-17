package main

import (
	"github.com/gorilla/mux"
	"github.com/jcunhasilva/golang-todo-list/task"
)

// MakeRouter builds the task routes
func MakeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/task", task.Index).Methods("GET")
	r.HandleFunc("/task", task.Create).Methods("POST")
	r.HandleFunc("/task", task.Update).Methods("PUT")
	r.HandleFunc("/task", task.Delete).Methods("DELETE")
	return r
}
