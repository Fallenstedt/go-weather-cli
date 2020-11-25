package services

import "time"

type IDateService interface {
	GetCurrentDate() string
}

type DateService struct {}

func NewDateService() *DateService {
	return &DateService{}
}

func (d DateService) GetCurrentDate() string {
	return time.Now().Format("Monday 15:04 Jan 02, 2006")
}
