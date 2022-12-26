package db

import "backend/api"


func (db *DBORM) AddMileageEarned(userid uint64, amountToAdd int, weekday string) error {
	item := api.MileageEarned{ID: userid}

	if err := db.Find(&item).Error; err != nil {
		return err
	}

	newAmount := item.Monday + amountToAdd
	return db.Model(&item).Update(weekday, newAmount).Error
}

func (db *DBORM) InitAllMileageEarned() error {
	return db.Model(&api.MileageEarned{}).Where("1 = 1").Select("*").Omit("user_id").Updates(api.MileageEarned{
		Monday: 0, Tuesday: 0, Wensday: 0, Thursday: 0, Friday: 0,
	}).Error
}
