package server

import (
	"log"
	"net/http"
	"time"

	"github.com/GoferAlex/go_final_project/pkg/api"
)

// структура сервера
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func Run(flog *log.Logger) *Server {

	api.Init()

	ServStandart := http.Server{
		Addr:         ":7540",
		Handler:      api.MyMux,
		ErrorLog:     flog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	Serv := Server{
		Logger: flog,
		Server: &ServStandart,
	}

	return &Serv
}
