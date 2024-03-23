package db

import (
	"gorm.io/gorm"
	"time"
)

type Advertisement struct {
	gorm.Model
	Title      string
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
	Conditions []Condition
}

type Condition struct {
	gorm.Model
	AdvertisementId uint
	AgeStart        int
	AgeEnd          int
	Gender          string
	Country         string
	Platform        string
}


