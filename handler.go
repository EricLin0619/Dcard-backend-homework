package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Advertisement
	_ = json.NewDecoder(r.Body).Decode(&post)
	byteSlice, _ := json.MarshalIndent(post, "", "  ")
	fmt.Println(string(byteSlice))
}

func GetAdHandler(w http.ResponseWriter, r *http.Request) {
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

