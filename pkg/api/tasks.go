package api

import (
	"net/http"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

const limit = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(limit) // в параметре максимальное количество записей
	if err != nil {
		wrong.Error = err.Error()
		writeJson(w, err)
		return
	}

	tasksResp := TasksResp{}
	tasksResp.Tasks = tasks

	writeJson(w, tasksResp)
}
