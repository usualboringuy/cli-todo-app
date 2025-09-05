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
	Status      bool
	CreatedOn   time.Time
	CompletedOn *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Status:      false,
		CompletedOn: nil,
		CreatedOn:   time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
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
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Status
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedOn = &completionTime
	}
	t[index].Status = !isCompleted
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
	table := table.New((os.Stdout))
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Status", "Created On", "Completed On")
	for index, t := range *todos {
		status := "❌"
		completedOn := ""

		if t.Status {
			status = "✔️"
			if t.CompletedOn != nil {
				completedOn = t.CompletedOn.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, status, t.CreatedOn.Format(time.RFC1123), completedOn)
	}
	table.Render()
}
