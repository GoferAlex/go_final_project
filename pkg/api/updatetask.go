package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	go func() {

		var task db.Task
		var wrong db.Wrong

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		tasksMu.Lock()
		defer tasksMu.Unlock()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				// читаем тело запроса
				body, err := io.ReadAll(r.Body)
				if err != nil {
					wrong.Error = err.Error()
					writeJson(w, wrong)
					return
				}
				defer r.Body.Close()

				// десериализуем JSON в Task
				if err := json.Unmarshal(body, &task); err != nil {
					wrong.Error = err.Error()
					writeJson(w, wrong)
					return
				}

				if task.Title == "" {
					err := errors.New("empty Title")
					wrong.Error = err.Error()
					writeJson(w, wrong)
					return
				}
				if err := checkDate(&task); err != nil {
					wrong.Error = err.Error()
					writeJson(w, wrong)
					return
				}

				err = db.UpdateTask(ctx, &task)
				if err != nil {
					wrong.Error = err.Error()
					writeJson(w, wrong)
					return
				}

				writeJson(w, task)

				return
			}
		}
	}()
}
