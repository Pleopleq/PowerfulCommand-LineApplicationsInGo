package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(index int) error {
	ls := *l

	if index <= 0 || index > len(ls) {
		return fmt.Errorf("Item %d does not exist", index)
	}
	ls[index-1].Done = true
	ls[index-1].CompletedAt = time.Now()

	return nil
}
