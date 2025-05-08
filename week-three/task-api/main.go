package main

import (
    "fmt"
    "net/http"
    "task-api/handler"
)

func main() {
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            handler.GetTasks(w, r)
        case "POST":
            handler.CreateTask(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
