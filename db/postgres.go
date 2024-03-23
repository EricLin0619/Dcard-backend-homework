package db
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewPostgresRepo() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Advertisement{}, &Condition{})
	return db
}