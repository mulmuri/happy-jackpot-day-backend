package db

import (
	"backend/api"
	"errors"
)



func (db *DBORM) TakeDailyMileage(userid uint64, daytype string) (int, error) {
	mileage := api.WeeklyMileage{ID: userid}
	if err := db.Model(mileage).Find(&mileage).Error; err != nil {
		return 0, err
	}

	if err:= db.Model(&api.WeeklyMileage{ID: userid}).Update(daytype, 0).Error; err != nil {
		return 0, err
	}

	switch (daytype) {
	case "monday":
		return mileage.Monday, nil
	case "tuesday":
		return mileage.Tuesday, nil
	case "wensday":
		return mileage.Wensday, nil
	case "thursday":
		return mileage.Tuesday, nil
	case "friday":
		return mileage.Friday, nil
	default:
		return 0, errors.New("invalid daytype")
	}
}


func (db *DBORM) AddDailyMileage(userid uint64, weekday string, amountToAdd int) error {
	item := api.WeeklyMileage{ID: userid}

	if err := db.Find(&item).Error; err != nil {
		return err
	}

	newAmount := item.Monday + amountToAdd
	return db.Model(&item).Update(weekday, newAmount).Error
}
