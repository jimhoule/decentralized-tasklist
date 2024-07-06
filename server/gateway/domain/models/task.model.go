package models

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsDone      bool   `json:"isDone"`
}