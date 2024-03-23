package app

import (
	"context"
	"github.com/EricLin0619/DcardBackend/handler"
	"github.com/gorilla/mux"
)

func (a *App) LoadRoutes() *mux.Router {
	ctx := context.Background()
	adHandler := &handler.Repo{Db: a.Db, Ctx: ctx}
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/ad", adHandler.CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", adHandler.GetAdHandler).Methods("GET")
	r.HandleFunc("/", adHandler.YourHandler).Methods("GET")
	return r
}
