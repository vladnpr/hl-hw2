package main

import "fmt"

type Task struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TasksList map[int]*Task

func (tl TasksList) Add(task *Task) {
	key := len(tl)
	_, ok := tl[key]
	if !ok {
		key = len(tl) + 1
	}

	tl[key] = task
	fmt.Println("Added task:", task.Title)
}

func (tl TasksList) Delete(key int) {
	_, ok := tl[key]
	if !ok {
		fmt.Printf("TaskList does not contain a Task with the specified id: %d n\n", key)
		return
	}

	delete(tl, key)
	fmt.Printf("Task %d has been deleted\n", key)

}

func (tl TasksList) Print() {
	if len(tl) == 0 {
		fmt.Println("TaskList is empty")
		return
	}

	fmt.Println("TaskList: ")
	for key, value := range tl {
		fmt.Printf("Task %d: %s\n", key, value.Title)
	}

}

func (tl TasksList) CompleteTask(taskID int) {
	_, ok := tl[taskID]

	if !ok {
		fmt.Printf("TaskList does not contain a Task with the specified id: %d n\n", key)
		return
	}

	task := tl[taskID]
	task.Completed = true
	fmt.Printf("Task %d marked as ompleted\n", taskID)
}

func main() {
	tasks := TasksList{}
	task0 := Task{Title: "Read a book", Completed: false}
	task1 := Task{Title: "Play some games", Completed: true}

	tasks.Add(&task0)
	tasks.Print()
	tasks.Add(&task1)
	tasks.Print()
	tasks.Delete(0)
	tasks.Print()
}
