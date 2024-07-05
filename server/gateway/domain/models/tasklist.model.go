package models

type Tasklist struct {
	Name        string
	Description string
	Email       string
	Tasks       []*Task
}