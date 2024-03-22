package application
import (
	"github.com/EricLin0619/DcardBackend/handler"
	"github.com/gorilla/mux"
)

func LoadRoutes () *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/ad", handler.CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", handler.GetAdHandler).Methods("GET")
	r.HandleFunc("/", handler.YourHandler)
	return r
}