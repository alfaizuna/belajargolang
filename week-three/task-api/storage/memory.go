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

func GetByID(id int) (model.Task, bool) {
    for _, t := range Tasks {
        if t.ID == id {
            return t, true
        }
    }
    return model.Task{}, false
}

func DeleteByID(id int) bool {
    for i, t := range Tasks {
        if t.ID == id {
            Tasks = append(Tasks[:i], Tasks[i+1:]...)
            return true
        }
    }
    return false
}