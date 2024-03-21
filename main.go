package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:		  0,  // 默认DB 0
	})
	ctx := context.Background()
	
	dbTest(rdb, ctx)
	
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)
	r.HandleFunc("/api/v1/ad", CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", GetAdHandler).Methods("GET")


    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}

func dbTest (db *redis.Client, ctx context.Context) {
	val, _ := db.Get(ctx, "test").Result()
	fmt.Println(val)
}





