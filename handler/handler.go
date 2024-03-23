package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type Repo struct {
	Db  *redis.Client
	Ctx context.Context
}

func (repo *Repo) CacheHandler(w http.ResponseWriter, r *http.Request) {
	var post Advertisement
	_ = json.NewDecoder(r.Body).Decode(&post)
	byteSlice, _ := json.MarshalIndent(post, "", "  ")
	repo.Db.RPush(repo.Ctx, "cacheList", string(byteSlice))             // push to list
	repo.Db.Publish(repo.Ctx, "cacheChannel", string(byteSlice)) // publish to channel
	// if err != nil {
	// 	panic(err)
	// }
	w.Write([]byte("push success!\n"))
}

func (repo *Repo) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Advertisement
	_ = json.NewDecoder(r.Body).Decode(&post)
	byteSlice, _ := json.MarshalIndent(post, "", "  ")
	fmt.Println(string(byteSlice))
}

func (repo *Repo) GetAdHandler(w http.ResponseWriter, r *http.Request) {
	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")
	age := r.URL.Query().Get("age")
	gender := r.URL.Query().Get("gender")
	country := r.URL.Query().Get("country")
	platform := r.URL.Query().Get("platform")
	fmt.Println("offset", offset)
	fmt.Println("limit", limit)
	fmt.Println("age", age)
	fmt.Println("gender", gender)
	fmt.Println("country", country)
	fmt.Println("platform", platform)
}

func (repo *Repo) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
	fmt.Print(repo.Db.Get(repo.Ctx, "test").Result())
}
