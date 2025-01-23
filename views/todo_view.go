package views

type TodoView struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTodoView(id int, title string, description string, completed bool) TodoView {
	return TodoView{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}
