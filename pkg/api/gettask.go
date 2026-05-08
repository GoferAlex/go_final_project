package api

import (
	"net/http"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	task, err := db.GetTask(id)

	if err != nil {
		writeJson(w, err)
	} else {
		writeJson(w, task)
	}
}
