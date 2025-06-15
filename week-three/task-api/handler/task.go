package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-api/model"
	"task-api/storage"

	"github.com/go-chi/chi/v5"
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

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Menggunakan chi.URLParam untuk mengambil parameter id
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, found := storage.GetByID(id)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Menggunakan chi.URLParam untuk mengambil parameter id
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if storage.DeleteByID(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

/*
Penjelasan perubahan untuk Chi Router:
1. Import diganti dari gorilla/mux ke github.com/go-chi/chi/v5
2. Pengambilan parameter URL menggunakan chi.URLParam() alih-alih mux.Vars()
3. Sintaks lebih sederhana dan lebih ringkas
4. Performa lebih baik dibanding gorilla/mux
5. Lebih mudah untuk menambahkan middleware jika diperlukan
*/
