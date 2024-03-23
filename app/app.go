package app

import (
	"log"
	"net/http"
	"github.com/EricLin0619/DcardBackend/db"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Port string
	Router *mux.Router
	Db *redis.Client
}

func NewApp(port string) *App {
	return &App{
		Port: ":"+port,
		Router: mux.NewRouter(),
		Db: db.NewDb(),
	}
}

func (a *App) Run() {
	a.Router = a.LoadRoutes()
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
	
}