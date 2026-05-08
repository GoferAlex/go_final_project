package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const formatDate = "20060102"

func afterNow(date, now time.Time) bool {
	if date.After(now) {
		return true
	} else {
		return false
	}
}

var ErrInvalidFormat = errors.New("invalid format of repeat")

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	date, err := time.Parse(formatDate, dstart)
	if err != nil {
		return "", err
	}

	rep := strings.Split(repeat, " ")

	d := 0
	switch len(rep) {

	case 1:
		if rep[0] != "y" {
			return "", ErrInvalidFormat
		} else {
			for {
				date = date.AddDate(1, 0, 0)
				if afterNow(date, now) {
					break
				}
			}
		}
	case 2:
		if rep[0] != "d" {
			return "", ErrInvalidFormat
		} else {
			d, err = strconv.Atoi(rep[1])
			if err != nil {
				return "", err
			}
			if d > 400 {
				return "", ErrInvalidFormat
			}
			for {
				date = date.AddDate(0, 0, d)
				if afterNow(date, now) {
					break
				}
			}
		}

	default:
		return "", ErrInvalidFormat
	}

	string := date.Format(formatDate)

	return string, nil
}

func nextDayHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("Request is not" + http.MethodGet)
		wrong.Error = err.Error()
		writeJson(res, wrong)
		return
	}

	now := req.FormValue("now")
	err := errors.New("")
	t := time.Now()

	if now != "" {
		t, err = time.Parse(formatDate, now)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}

	date := req.FormValue("date")

	repeat := req.FormValue("repeat")
	answer, err := NextDate(t, date, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Write([]byte(answer))
}
