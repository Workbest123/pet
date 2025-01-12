package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// go run todolist.go -add "Buy products"
// go run todolist.go -list
// go run todolist.go -delete 1
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const taskFile = "tasks.json"

// loadTasks загружает задачи из JSON файла
func loadTasks() (TaskList, error) {
	var taskList TaskList

	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return taskList, nil
	}

	data, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return taskList, fmt.Errorf("failed to read tasks file: %w", err)
	}

	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return taskList, fmt.Errorf("failed to parse tasks file: %w", err)
	}

	return taskList, nil
}

// saveTasks сохраняет задачи в JSON файл
func saveTasks(taskList TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize tasks: %w", err)
	}

	err = ioutil.WriteFile(taskFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}

	return nil
}

// addTask добавляет новую задачу
func addTask(name string) error {
	taskList, err := loadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:   len(taskList.Tasks) + 1,
		Name: name,
	}
	taskList.Tasks = append(taskList.Tasks, newTask)

	return saveTasks(taskList)
}

// listTasks выводит все задачи
func listTasks() error {
	taskList, err := loadTasks()
	if err != nil {
		return err
	}

	if len(taskList.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return nil
	}

	for _, task := range taskList.Tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}

	return nil
}

// deleteTask удаляет задачу по ID
func deleteTask(id int) error {
	taskList, err := loadTasks()
	if err != nil {
		return err
	}

	updatedTasks := make([]Task, 0)
	for _, task := range taskList.Tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(taskList.Tasks) {
		return fmt.Errorf("task with ID %d not found", id)
	}

	taskList.Tasks = updatedTasks
	return saveTasks(taskList)
}

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	delete := flag.Int("delete", 0, "Delete a task by ID")
	flag.Parse()

	switch {
	case *add != "":
		err := addTask(*add)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added successfully.")

	case *list:
		err := listTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
		}

	case *delete != 0:
		err := deleteTask(*delete)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
		fmt.Println("Task deleted successfully.")

	default:
		fmt.Println("Usage:")
		fmt.Println("  -add <task name>    Add a new task")
		fmt.Println("  -list               List all tasks")
		fmt.Println("  -delete <task ID>   Delete a task by ID")
	}
}
