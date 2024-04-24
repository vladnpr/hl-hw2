package main

import "fmt"

type Task struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TasksList map[int]*Task

func main() {
	tasks := TasksList{}
	task0 := Task{Title: "Read a book", Completed: false}
	task1 := Task{Title: "Play some games", Completed: true}

	tasks.add(&task0)
	tasks.print()
	tasks.add(&task1)
	tasks.print()
	tasks.delete(0)
	tasks.print()
}

func (tl *TasksList) add(task *Task) {
	var key int

	for {
		_, ok := (*tl)[key]
		if !ok {
			(*tl)[key] = task
			fmt.Println("Added task:", task.Title)
			break
		} else {
			key++
			continue
		}
	}
}

func (tl *TasksList) delete(key int) {
	_, ok := (*tl)[key]
	if !ok {
		fmt.Printf("TaskList do not contains Task with id: %d n\n", key)
	} else {
		delete(*tl, key)
		fmt.Printf("Task %d has been deleted\n", key)
	}
}

func (tl *TasksList) print() {
	if len(*tl) == 0 {
		fmt.Println("TaskList is empty")
	} else {
		fmt.Println("TaskList: ")
		for key, value := range *tl {
			fmt.Printf("Task %d: %s\n", key, value.Title)
		}
	}
}
