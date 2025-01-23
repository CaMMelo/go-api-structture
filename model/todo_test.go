package model

import "testing"

func TestToggle(t *testing.T) {
	todo := &Todo{
		ID:          1,
		Title:       "",
		Description: "",
		Completed:   true,
	}
	todo.Toggle()
	if todo.Completed != false {
		t.Fatalf("did not toggle")
	}
	todo.Toggle()
	if todo.Completed != true {
		t.Fatalf("did not toggle")
	}
}
