package models

type Tasklist struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Email       string  `json:"email"`
	Tasks       []*Task `json:"tasks,omitempty"`
}