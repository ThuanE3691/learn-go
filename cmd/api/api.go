package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ThuanE3691/learn-go/cmd/server/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: 	addr,
		db: 		db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router);
}