package api

import (
	"context"
	"net/http"
	"time"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	go func() {

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		wrong := db.Wrong{
			Error: "задача не найдена",
		}

		tasksMu.RLock()
		defer tasksMu.RUnlock()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				id := r.URL.Query().Get("id")
				task, err := db.GetTask(ctx, id)

				if err != nil {
					writeJson(w, wrong)
					return
				}

				writeJson(w, task)

				return
			}
		}
	}()
}
