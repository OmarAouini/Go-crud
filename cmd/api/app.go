package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//App type
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//initialize the App
func (a *App) Initialize(user, password, dbname string) {}

//run the app given a port
func (a *App) Run(addr string) {}
