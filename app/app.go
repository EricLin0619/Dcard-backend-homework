package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/EricLin0619/DcardBackend/db"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)
type Advertisement struct {
	Title string `json:"title"`
	StartAt time.Time `json:"startAt"`
	EndAt time.Time `json:"endAt"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	AgeStart *int `json:"ageStart"`
	AgeEnd *int `json:"ageEnd"`
	Gender []string `json:"gender"`
	Country []string `json:"country"`
	Platform []string `json:"platform"`
}

type App struct {
	Port       string
	Router     *mux.Router
	RedisDb    *redis.Client
	PostgresDb *gorm.DB
}

func NewApp(port string) *App {
	return &App{
		Port:       ":" + port,
		Router:     mux.NewRouter(),
		RedisDb:    db.NewRedisRepo(),
		PostgresDb: db.NewPostgresRepo(),
	}
}

func (a *App) Run() {
	a.Router = a.LoadRoutes()
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}

func (a *App) ProcessWriteQueue() {
	ctx := context.Background()
	var adInRedis Advertisement
	data, _ := a.RedisDb.LPop(ctx, "cacheList").Result()
	json.Unmarshal([]byte(data), &adInRedis) // unmarshal to struct
	fmt.Println(reflect.TypeOf(adInRedis.StartAt))
	fmt.Println(adInRedis.StartAt)

	var advertisementId int
	a.PostgresDb.Exec("INSERT INTO advertisements (title, start_at, end_at) VALUES ($1, $2, $3)", adInRedis.Title, adInRedis.StartAt, adInRedis.EndAt)
	a.PostgresDb.Raw("SELECT max(id) FROM advertisements").Scan(&advertisementId)
	for _, condition := range adInRedis.Conditions {
		gender := strings.Join(condition.Gender, ", ")
		country := strings.Join(condition.Country, ", ")
		platform := strings.Join(condition.Platform, ", ")
		a.PostgresDb.Exec("INSERT INTO conditions (advertisement_id, age_start, age_end, gender, country, platform) VALUES ($1, $2, $3, $4, $5, $6)", advertisementId, condition.AgeStart, condition.AgeEnd, gender, country, platform)
	}
	fmt.Println("Success")
}

// func (a *App) ListenToChannel() {
// 	fmt.Println("ListenToChannel goroutine started")
// 	ctx := context.Background()
// 	pubsub := a.RedisDb.Subscribe(ctx, "cacheChannel")
// 	for {
// 		msg, err := pubsub.ReceiveMessage(ctx)
// 		if err != nil {
// 			fmt.Println("There is an error in ListenToChannel goroutine.")
// 			panic(err)
// 		}
// 		fmt.Println(msg.Channel, msg.Payload)
// 	}
// }

// func (a *App) ConnectPostgres() {
// 	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db.AutoMigrate(&User{}, &Card{})
// 	// db.Create(&User{Name: "coco", Age: 25, Cards: []Card{{Code: "1234"}, {Code: "5678"}}})

// 	var target User
// 	db.Preload("Cards").Find(&target, 10)
// 	fmt.Println(target.Name)
// 	fmt.Println(target.Cards[0].Code)
// }
