package main

import (
    "fmt"
    "net/http"
    "task-api/handler"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", handler.GetTaskByID).Methods("GET")
    r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
