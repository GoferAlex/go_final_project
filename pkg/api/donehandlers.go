package api

import (
	"fmt"
	"net/http"
	"time"

	_ "modernc.org/sqlite"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

type Empty struct{}

var (
	empty    Empty
	nextDate string
	wrong    respErr
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("Request is not" + http.MethodPost)
		writeJson(w, err)
		return
	}

	id := r.URL.Query().Get("id")

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, err)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJson(w, err)
			return
		}
		writeJson(w, empty)
		return
	}
	nextDate, err = NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJson(w, err)
		return
	}

	err = db.UpdateDate(nextDate, id)
	if err != nil {
		writeJson(w, err)
		return
	}

	writeJson(w, empty)
}
