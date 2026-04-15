package server

import (
	"net/http"
)

func main() {

    webDir := "github.com/GoferAlex/go_final_project/web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
}