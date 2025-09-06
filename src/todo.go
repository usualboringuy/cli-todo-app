package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
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
	tab := table.New(os.Stdout)
	tab.SetPadding(2)
	tab.SetHeaderStyle(table.StyleBold)
	tab.SetLineStyle(table.StyleBrightBlue)
	tab.SetDividers(table.UnicodeRoundedDividers)
	tab.SetAlignment(table.AlignCenter, table.AlignLeft, table.AlignCenter, table.AlignCenter, table.AlignCenter)
	tab.SetRowLines(false)
	tab.SetHeaders("#", "Title", "Status", "Created On", "Completed On")
	for index, t := range *todos {
		status := "Not Done"
		completedOn := ""
		if t.Status {
			status = "Done"
			if t.CompletedOn != nil {
				completedOn = t.CompletedOn.Format("Jan 02 15:04")
				tab.AddRow(
					tml.Sprintf("<green>%s</green>", strconv.Itoa(index)),
					t.Title,
					tml.Sprintf("<green>%s</green>", status),
					tml.Sprintf("<green>%s</green>", t.CreatedOn.Format("Jan 02 15:04")),
					tml.Sprintf("<green>%s</green>", completedOn),
				)
				continue
			}
		}
		tab.AddRow(
			tml.Sprintf("<red>%s</red>", strconv.Itoa(index)),
			t.Title,
			tml.Sprintf("<red>%s</red>", status),
			tml.Sprintf("<red>%s</red>", t.CreatedOn.Format("Jan 02 15:04")),
			tml.Sprintf("<red>%s</red>", completedOn),
		)
	}
	tab.Render()
}
