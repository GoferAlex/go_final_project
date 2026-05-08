package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	// читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJson(w, err)
		return
	}
	defer r.Body.Close()

	// десериализуем JSON в Task
	if err = json.Unmarshal(body, &task); err != nil {
		writeJson(w, err)
		return
	}

	if task.Title == "" {
		err = errors.New("empty Title")
		writeJson(w, err)
		return
	}

	if err = checkDate(&task); err != nil {
		writeJson(w, err)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(w, err)
		return
	}

	writeJson(w, task)
}
