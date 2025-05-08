package storage

import (
    "task-api/model"
)

var Tasks []model.Task
var NextID = 1

func GetAll() []model.Task {
    return Tasks
}

func Add(task model.Task) model.Task {
    task.ID = NextID
    NextID++
    Tasks = append(Tasks, task)
    return task
}
