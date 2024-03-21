package main
import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type Advertisement struct {
	Title string `json:"title"`
	StartAt time.Time `json:"start_at"`
	EndAt time.Time `json:"end_at"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	AgeStart *int `json:"ageStart"`
	AgeEnd *int `json:"ageEnd"`
	Gender []string `json:"gender"`
	Country []string `json:"country"`
	Platform []string `json:"platform"`
}


func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)
	r.HandleFunc("/api/v1/ad", CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", GetAdHandler).Methods("GET")


    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}




