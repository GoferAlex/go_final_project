package api

import (
	"net/http"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	var wrong db.Wrong

	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		wrong.Error = err.Error()
		writeJson(w, wrong)
		return
	}

	tasksResp := TasksResp{}
	tasksResp.Tasks = tasks

	writeJson(w, tasksResp)
}
