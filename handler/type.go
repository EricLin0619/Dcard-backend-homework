package handler

import "time"


type Advertisement struct {
	Title string `json:"title"`
	StartAt *time.Time `json:"startAt"`
	EndAt *time.Time `json:"endAt"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	AgeStart *int `json:"ageStart"`
	AgeEnd *int `json:"ageEnd"`
	Gender []string `json:"gender"`
	Country []string `json:"country"`
	Platform []string `json:"platform"`
}