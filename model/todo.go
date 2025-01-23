package model

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func (todo *Todo) Toggle() {
	todo.Completed = !todo.Completed
}
