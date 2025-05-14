package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := (*todos) // Dereferencing to get the slice
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// Assuming TodoItem has a 'Completed' boolean field
	// and 'CompletedAt' a *time.Time field.
	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	} else {
		// Optionally, clear the CompletedAt time if toggling back to not completed
		t[index].CompletedAt = nil
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	// Assuming 'table' refers to a package like tablewriter.
	// For example: table := tablewriter.NewWriter(os.Stdout)
	table := table.New(os.Stdout) // Using tablewriter as an example
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		// Assuming t.Title is a string and t.CreatedAt is a time.Time
		// (not a pointer for this direct format call)
		table.AddRow(
			strconv.Itoa(index), // Typically, display index as 1-based
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	table.Render()
}
