package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var shema = `
    CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(256) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(256) NOT NULL DEFAULT ""
);

CREATE INDEX task_date ON scheduler (date);
`
var Datbase *sql.DB

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	/*db*/
	Datbase, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer Datbase.Close()

	if install {
		_, err = Datbase.Exec(shema)
		if err != nil {
			return err
		}
	}
	return nil
}
