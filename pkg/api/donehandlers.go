package api

import (
	"context"
	"net/http"
	"time"

	_ "modernc.org/sqlite"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

type Empty struct{}

var (
	empty    Empty
	nextDate string
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wrong := db.Wrong{
		Error: "Error task done",
	}

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

			if task.Repeat == "" {
				err = db.DeleteTask(ctx, id)
				if err != nil {
					writeJson(w, wrong)
					return
				}
				writeJson(w, empty)
				return
			}
			nextDate, err = NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				writeJson(w, wrong)
				return
			}

			err = db.UpdateDate(ctx, nextDate, id)
			if err != nil {
				writeJson(w, wrong)
				return
			}

			writeJson(w, empty)

			return
		}
	}
}
