package api

import (
	"net/http"
	"sync"
)

var (
	MyMux   = http.NewServeMux()
	tasksMu sync.RWMutex
)

func Init() {
	webDir := "./web"

	MyMux.Handle("/", http.FileServer(http.Dir(webDir)))
	MyMux.HandleFunc("/api/nextdate", nextDayHandler)
	MyMux.HandleFunc("/api/task", taskHandler)
	MyMux.HandleFunc("/api/tasks", tasksHandler)
	MyMux.HandleFunc("/api/task/done", taskDoneHandler)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// обработка метода Post
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	}
}
