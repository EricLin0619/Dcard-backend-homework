package db

import (
	"github.com/redis/go-redis/v9"
	"context"
	"fmt"
	// "fmt"
)

// var ctx = context.Background()

type Database struct {
	Db *redis.Client
	Ctx context.Context
}

func NewDb () *Database{
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", 
		DB:		  0,  
	})
	db := Database {
		Db: rdb,
		Ctx: ctx,
	}
	return &db
}

func (db *Database) DbTest()  {
	val, _ := db.Db.Get(db.Ctx, "test").Result()
	fmt.Println(val)
}
