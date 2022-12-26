package api

import (
	"errors"
)



type MileageEarned struct {
	ID       uint64 `gorm:"column:user_id"  json:"-"`
	Key      string `gorm:"-"               json:"ID"`
	Monday   int    `gorm:"column:monday"   json:"monday"`
	Tuesday  int    `gorm:"column:tuesday"  json:"tuesday"`
	Wensday  int    `gorm:"column:wensday"  json:"wensday"`
	Thursday int    `gorm:"column:thursday" json:"thursday"`
	Friday   int    `gorm:"column:friday"   json:"friday"`
}

func (MileageEarned) TableName() string {
	return "mileage_earned"
}



type WeeklyMileage struct {
	ID       uint64 `gorm:"column:user_id"  json:"-"`
	Key      string `gorm:"-"               json:"ID"`
	Monday   int    `gorm:"column:monday"   json:"monday"`
	Tuesday  int    `gorm:"column:tuesday"  json:"tuesday"`
	Wensday  int    `gorm:"column:wensday"  json:"wensday"`
	Thursday int    `gorm:"column:thursday" json:"thursday"`
	Friday   int    `gorm:"column:friday"   json:"friday"`
}

func (WeeklyMileage) TableName() string {
	return "weekly_mileage"
}

func (u *WeeklyMileage) GetSum() int {
	return u.Monday + u.Tuesday + u.Wensday + u.Thursday + u.Friday
}

var Monday string = "Monday"
var Tuesday string = "Tuesday"
var Wensday string = "Wensday"
var Thursday string = "Thursday"
var Friday string = "Friday"

var DayOfTheWeek = [5]string{Monday, Tuesday, Wensday, Thursday, Friday}

type DailyMileageRequest struct {
	ID      uint64
	Key     string `json:"userID"`
	Amount  int    `json:"amount"`
	Weekday string `json:"weekday"`
	State   string `json:"state"`
}

func WeekdayValid(weekday string) error {
	for _, day := range DayOfTheWeek {
		if day == weekday {
			return nil
		}
	}

	return errors.New(DailyMileageReqStatusInvalidInput)
}

var DailyMileageReqStatusValid string = "Valid"
var DailyMileageReqStatusInvalidInput string = "Invalid Amount"
var DailyMileageReqStatusUserNotFound string = "User Not Found"
