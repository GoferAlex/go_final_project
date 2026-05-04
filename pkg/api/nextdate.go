package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const api = "20060102"

func afterNow(now, date time.Time) bool {
	if date.After(now) {
		return true
	}
	return false
}

var ErrInvalidFormat = errors.New("invalid format of repeat")

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(api, dstart)
	if err != nil {
		return "", err
	}

	rep := strings.Split(repeat, " ")

	d := 0
	if len(rep) == 1 && rep[0] != "y" {
		return "", ErrInvalidFormat
	}
	if len(rep) == 2 {
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
		}
	}
	if len(rep) > 2 {
		return "", ErrInvalidFormat
	}

	if rep[0] == "y" {
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	} else {
		for {
			date = date.AddDate(0, 0, d)
			if afterNow(date, now) {
				break
			}
		}
	}
	string := date.Format(api)

	return string, nil
}

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	err := errors.New("")
	t := time.Now()
	if now != "" {
		t, err = time.Parse(api, now)
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
