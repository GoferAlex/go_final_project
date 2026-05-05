package api

import (
	"context"
	"net/http"
	"time"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	go func() {

		var wrong db.Wrong

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		tasksMu.RLock()
		defer tasksMu.RUnlock()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				tasks, err := db.Tasks(ctx, 50) // в параметре максимальное количество записей
				if err != nil {
					wrong.Error = err.Error()
					writeJson(w, wrong)
					return
				}

				tasksResp := TasksResp{}
				tasksResp.Tasks = tasks

				writeJson(w, tasksResp)

				return
			}
		}
	}()
}
