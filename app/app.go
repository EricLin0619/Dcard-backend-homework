package app

import (
	"log"
	"net/http"
	"context"
	"github.com/EricLin0619/DcardBackend/db"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"fmt"
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

func (a *App) ListenToChannel() {
	fmt.Println("ListenToChannel goroutine started")
	ctx := context.Background()
	pubsub := a.Db.Subscribe(ctx, "cacheChannel")
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("There is an error in ListenToChannel goroutine.")
			panic(err)
		}
		fmt.Println(msg.Channel, msg.Payload)
	}
}