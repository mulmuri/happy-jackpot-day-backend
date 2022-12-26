package db

import (
	"backend/api"
)



func (db *DBORM) InitAllWeeklyMileage() error {
	return db.Model(&api.WeeklyMileage{}).Where("1 = 1").Select("*").Omit("user_id").Updates(api.WeeklyMileage{
		Monday: 0, Tuesday: 0, Wensday: 0, Thursday: 0, Friday: 0,
	}).Error
}

func (db *DBORM) GetWeeklyMileage(userid uint64) (api.WeeklyMileage, error) {
	mileage := api.WeeklyMileage{ID: userid}
	return mileage, db.Find(&mileage).Error
}