package api

import (
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

	wrong := db.Wrong{
		Error: "Error task done",
	}

	id := r.URL.Query().Get("id")
	task, err := db.GetTask(id)

	if err != nil {
		writeJson(w, wrong)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
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

	err = db.UpdateDate(nextDate, id)
	if err != nil {
		writeJson(w, wrong)
		return
	}

	writeJson(w, empty)
}
