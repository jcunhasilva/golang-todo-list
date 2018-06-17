package task

import (
	"log"
	"net/http"

	"github.com/jcunhasilva/golang-todo-list/service"
)

//Index get the list of tasks
func Index(w http.ResponseWriter, r *http.Request) {
	tasks, err := getAllTasks()
	if err != nil {
		log.Printf("Failed to get the list of tasks: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service.SendJSONResponse(w, tasks, http.StatusOK)
}

//Create new task
func Create(w http.ResponseWriter, r *http.Request) {
	var msg CreateMessage
	service.ReceiveAsJSON(r, &msg)
	task, err := msg.newTask()
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.SendJSONResponse(w, task, http.StatusCreated)
}

//Update close a task
func Update(w http.ResponseWriter, r *http.Request) {
	var msg UpdateMessage
	service.ReceiveAsJSON(r, &msg)
	task, err := msg.closeTask()
	if err != nil {
		log.Printf("Failed to close task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.SendJSONResponse(w, task, http.StatusOK)
}

//Delete cancel a task
func Delete(w http.ResponseWriter, r *http.Request) {
	var msg UpdateMessage
	service.ReceiveAsJSON(r, &msg)
	if err := msg.deleteTask(); err != nil {
		log.Printf("Failed to delete task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.SendJSONResponse(w, msg, http.StatusOK)
}
