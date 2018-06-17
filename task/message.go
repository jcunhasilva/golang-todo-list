package task

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//CreateMessage Payload to create a task
type CreateMessage struct {
	Name    string    `json:"name" validate:"required"`
	DueDate time.Time `json:"due_date" validate:"required"`
}

//create task from message
func (msg CreateMessage) newTask() (*Task, error) {
	if err := validate.Struct(msg); err != nil {
		return nil, err
	}

	task, err := createTask(msg.Name, msg.DueDate)
	return task, err
}

//UpdateMessage Payload to close a task
type UpdateMessage struct {
	Name string `json:"name" validate:"required"`
}

func (msg UpdateMessage) closeTask() (*Task, error) {
	if err := validate.Struct(msg); err != nil {
		return nil, err
	}
	task, err := closeTask(msg.Name)
	return task, err
}

func (msg UpdateMessage) deleteTask() error {
	if err := validate.Struct(msg); err != nil {
		return err
	}
	return deleteTask(msg.Name)
}
