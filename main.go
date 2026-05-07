package main

import (
	"fmt"

	"github.com/GoferAlex/go_final_project/pkg/db"
	"github.com/GoferAlex/go_final_project/pkg/server"
)

func main() {
	//create and open the database
	err := db.Init("scheduler.db")
	if err != nil {
		fmt.Println("Run database error", err)
	}

	server.Run()
}
