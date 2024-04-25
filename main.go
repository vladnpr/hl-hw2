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
	if ok {
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
		fmt.Println("\r\rTaskList is empty")
		return
	}

	fmt.Println("TaskList: ")
	for key, value := range tl {
		fmt.Printf("TaskID: %d Title: %s Completed: %t \n", key, value.Title, value.Completed)
	}

}

func (tl TasksList) CompleteTask(taskID int) {
	_, ok := tl[taskID]

	if !ok {
		fmt.Printf("TaskList does not contain a Task with the specified id: %d n\n", taskID)
		return
	}

	task := tl[taskID]
	task.Completed = true
	fmt.Printf("Task %d marked as Completed\n", taskID)
}

func main() {
	tasks := make(TasksList)

	for {
		fmt.Println("\n1. Add a task")
		fmt.Println("2. Remove a task")
		fmt.Println("3. View all tasks")
		fmt.Println("4. Mark a task as completed")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			var title string
			fmt.Print("Enter task title: ")
			if _, err := fmt.Scanln(&title); err != nil {
				fmt.Println("Error: Failed to read task title.")
				continue
			}

			task := Task{Title: title}
			tasks.Add(&task)
		case 2:
			var id int
			fmt.Print("Enter ID of the task to remove: ")
			if _, err := fmt.Scanln(&id); err != nil {
				fmt.Println("Error: Failed to read task ID.")
				continue
			}

			tasks.Delete(id)
		case 3:
			tasks.Print()
		case 4:
			var id int
			fmt.Print("Enter ID of the task to mark as completed: ")
			if _, err := fmt.Scanln(&id); err != nil {
				fmt.Println("Error: Failed to read task ID.")
				continue
			}

			tasks.CompleteTask(id)
		case 5:
			fmt.Println("Program exited.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
