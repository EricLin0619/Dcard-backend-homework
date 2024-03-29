package app
import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/EricLin0619/DcardBackend/utils"
	"github.com/EricLin0619/DcardBackend/db"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

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

type Result struct {
	Title string `json:"title"`
	StartAt time.Time `json:"startAt"`
	EndAt time.Time `json:"endAt"`
	AgeStart *int `json:"age_start"`
	AgeEnd *int `json:"age_end"`
}

func (a *App) GetDataTest()  {
	var ad []Result
	a.PostgresDb.Model(&db.Advertisement{}).Select("title, start_at, end_at, conditions.age_start, conditions.age_end").Joins("JOIN conditions ON advertisements.id = conditions.advertisement_id").Scan(&ad)
	fmt.Println(utils.StructToJson(ad))
}

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
