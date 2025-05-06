package main

import "fmt"

type Task struct {
	Title string
	Done bool
}

func main() {
	tasks := []Task{
		{Title: "Belajar Go", Done: false},
	    {Title: "Bikin CLI", Done: true},
	}

	for i, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", i+1, task.Title, status)
	}
}