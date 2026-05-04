package server

import (
	"fmt"
	"net/http"

	"github.com/GoferAlex/go_final_project/pkg/api"
)

func Run() {

	api.Init()

	// run server
	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		fmt.Println(err)
	}
}
