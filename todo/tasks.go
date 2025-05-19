package todo

import "time"

type Task struct {
	Name   string
	Text   string
	IsDone bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func (t *Task) Done() {
	doneTime := time.Now()

	t.IsDone = true
	t.DoneAt = &doneTime
}

func NewTask(name string, text string) *Task {
	return &Task{
		Name:   name,
		Text:   text,
		IsDone: false,

		CreatedAt: time.Now(),
		DoneAt:    nil,
	}
}
