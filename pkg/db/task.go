package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type Wrong struct {
	Error string `json:"error"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat)` +
		`VALUES (:date, :title, :comment, :repeat)`
	res, err := Datbase.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return id, err
	}
	r, err := res.LastInsertId()
	if err != nil {
		return id, err
	}
	return r, nil
}

func Tasks(limit int) ([]*Task, error) {
	tasks := []*Task{}
	rows, err := Datbase.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date")
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {

	task := Task{}

	err := Datbase.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return &task, err
	}

	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat` +
		`WHERE id = :id`
	res, err := Datbase.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}

	return nil
}

func DeleteTask(id string) error {

	_, err := Datbase.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}

func UpdateDate(next, id string) error {
	res, err := Datbase.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
		sql.Named("date", next),
		sql.Named("id", id))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating date`)
	}

	return nil
}
