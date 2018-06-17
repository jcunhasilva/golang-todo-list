package task

import (
	"time"

	"github.com/jcunhasilva/golang-todo-list/config"
	"github.com/jcunhasilva/golang-todo-list/model"
)

// Status is the task status
type Status int

const (
	//TaskOpen new task
	TaskOpen Status = iota + 1
	//TaskClosed finished task
	TaskClosed
)

//Task entity
type Task struct {
	model.BaseModel
	ID      int
	Name    string    `sql:",notnull"`
	DueDate time.Time `sql:",notnull"`
	Status  int
}

func getAllTasks() (tasks []*Task, err error) {
	err = config.DBConnection.Client.Model(&tasks).Select()
	return
}

func createTask(name string, dueDate time.Time) (*Task, error) {
	task := &Task{
		Name:    name,
		DueDate: dueDate,
		Status:  int(TaskOpen),
	}

	if err := config.DBConnection.Client.Insert(task); err != nil {
		return nil, err
	}

	return task, nil
}

func closeTask(name string) (*Task, error) {
	task := new(Task)
	if err := config.DBConnection.Client.Model(task).Where("name = ?", name).Select(); err != nil {
		return nil, err
	}
	_, err := config.DBConnection.Client.Model(task).Set("status = ?, updated_at = ?", int(TaskClosed), time.Now()).Where("id = ?id").Update()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func deleteTask(name string) error {
	task := new(Task)
	_, err := config.DBConnection.Client.Model(task).Where("name = ?", name).Delete()
	return err
}
