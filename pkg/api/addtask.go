package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	var wrong db.Wrong

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

	id, err := db.AddTask(&task)
	if err != nil {
		wrong.Error = err.Error()
		writeJson(w, wrong)
		return
	}

	task.ID = strconv.Itoa(int(id))
	writeJson(w, task)
}

func checkDate(task *db.Task) error {
	now := time.Now().Format(api)
	timeNow, err := time.Parse(api, now)
	if err != nil {
		return err
	}

	var (
		date time.Time
		next string
	)

	if task.Date == "" {
		task.Date = now
	} else {
		date, err = time.Parse(api, task.Date)
		if err != nil {
			return err
		}
		if afterNow(timeNow, date) {
			if len(task.Repeat) == 0 {
				// если правила повторения нет, то берём сегодняшнее число
				task.Date = now
			} else {
				// в противном случае, берём вычисленную ранее следующую дату
				next, err = NextDate(timeNow, task.Date, task.Repeat)
				if err != nil {
					return err
				}
				task.Date = next
			}
		}
	}
	return nil
}

func writeJson(w http.ResponseWriter, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
