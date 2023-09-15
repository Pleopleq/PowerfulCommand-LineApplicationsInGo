package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) TaskPrintFormatter(verbose bool) string {
	formatted := ""

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)

		if verbose {
			formatted += fmt.Sprintf(" - %s \n", t.CreatedAt.Format(time.RFC822))
		}
	}

	return formatted
}

func (l *List) String() string {
	formatted := l.TaskPrintFormatter(false)
	return formatted
}

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

	fmt.Print(l)

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

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) NotComplete() {
	formatted := ""

	for k, t := range *l {
		if !t.Done {
			formatted += fmt.Sprintf("%d: %s\n", k+1, t.Task)
		}
	}

	fmt.Println(formatted)
}
