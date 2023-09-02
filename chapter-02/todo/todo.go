package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (l *List) Delete(index int) error {
	ls := *l

	if index <= 0 || index > len(ls) {
		return fmt.Errorf("Item %d does not exist", index)
	}

	*l = append(ls[:index-1], ls[index:]...)

	return nil
}

func (l *List) Save(filename string) error {
	json, err := json.Marshal(l)

	if err != nil {
		return err
	}

	//IOUTIL IS DEPRECATED
	return ioutil.WriteFile(filename, json, 0644)
}
