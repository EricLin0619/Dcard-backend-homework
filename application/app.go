package application
import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type App struct {
	Port string
	Router *mux.Router
}

func NewApp(port string) *App {
	return &App{
		Port: ":"+port,
		Router: LoadRoutes(),
	}
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}