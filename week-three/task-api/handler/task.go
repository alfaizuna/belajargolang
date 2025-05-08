package handler

import (
    "encoding/json"
    "net/http"
    "task-api/model"
    "task-api/storage"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(storage.GetAll())
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newTask := storage.Add(task)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newTask)
}
