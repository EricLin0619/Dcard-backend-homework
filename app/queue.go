package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EricLin0619/DcardBackend/handler"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
	"time"
)

func (a *App) ProcessWriteQueue() {
	// Lpop from redis cache list
	ctx := context.Background()
	var adInRedis handler.Advertisement
	var advertisementId int
	
	for {
		data, err := a.RedisDb.LPop(ctx, "cacheList").Result()
		if err != nil {
			if err == redis.Nil {
				time.Sleep(1 * time.Second)
				continue
			} else {
				log.Println("Error dequeuing write operation:", err)
				continue
			}
		}
		json.Unmarshal([]byte(data), &adInRedis) // unmarshal to struct
		
		a.PostgresDb.Exec("INSERT INTO advertisements (title, start_at, end_at) VALUES ($1, $2, $3)", adInRedis.Title, adInRedis.StartAt, adInRedis.EndAt)
		a.PostgresDb.Raw("SELECT max(id) FROM advertisements").Scan(&advertisementId)
		for _, condition := range adInRedis.Conditions {
			gender := strings.Join(condition.Gender, ", ")
			country := strings.Join(condition.Country, ", ")
			platform := strings.Join(condition.Platform, ", ")
			a.PostgresDb.Exec("INSERT INTO conditions (advertisement_id, age_start, age_end, gender, country, platform) VALUES ($1, $2, $3, $4, $5, $6)", advertisementId, condition.AgeStart, condition.AgeEnd, gender, country, platform)
		}
		fmt.Println("Insert to postgres Success")
	}

}
