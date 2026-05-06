package api

import (
	"fmt"
	"net/http"

	"github.com/GoferAlex/go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dTH go")

	wrong := db.Wrong{
		Error: "Error delete",
	}

	id := r.URL.Query().Get("id")
	_, err := db.GetTask(id)
	if err != nil {
		writeJson(w, wrong)
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		writeJson(w, wrong)
		return
	}

	writeJson(w, empty)
}
