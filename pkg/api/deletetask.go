package api

import (
	"context"
	"net/http"
	"time"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	go func() {

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		wrong := db.Wrong{
			Error: "Error delete",
		}

		tasksMu.Lock()
		defer tasksMu.Unlock()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				id := r.URL.Query().Get("id")
				_, err := db.GetTask(ctx, id)

				if err != nil {
					writeJson(w, wrong)
					return
				}

				err = db.DeleteTask(ctx, id)

				if err != nil {
					writeJson(w, wrong)
					return
				}

				writeJson(w, empty)

				return
			}
		}
	}()
}
