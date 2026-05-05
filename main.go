package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GoferAlex/go_final_project/pkg/db"
	"github.com/GoferAlex/go_final_project/pkg/server"
)

func main() {
	// create and open loggers file
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// close loggers file in the end
	defer file.Close()

	// create logger, which we will use
	mylog := log.New(file, "", log.LstdFlags|log.Lshortfile)

	//create and open the database
	err = db.Init("scheduler.db")
	if err != nil {
		mylog.Print(err)
		fmt.Println("Run database error", err)
	}
	defer db.Datbase.Close()

	// create server
	s := server.Run(mylog)

	// run server
	err = http.ListenAndServe(s.Server.Addr, s.Server.Handler)
	if err != nil {
		mylog.Print(err)
		mylog.Fatal(err)
	}
}
