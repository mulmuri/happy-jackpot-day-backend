package db

import (
	"backend/api"
)

func (db *DBORM) AddDailyMileageByDegree(userid uint64, degree, amountToAdd int) error {
	item := api.MileageByDegree{
		ID: userid,
		Degree: degree,
	}

	result := db.Find(&item)
	
	if result.RowsAffected == 0 {
		item.Amount = 0
		return db.Create(&item).Error;
	}

	newAmount := item.Amount + amountToAdd
	return db.Model(&item).Update("amoount", newAmount).Error
}

func (db *DBORM) GetAllMileageByDegree(userid uint64) ([]api.MileageByDegree, error) {
	items := []api.MileageByDegree{}
	return items, db.Find(&items, "user_id = ?", userid).Error;
}

func (db *DBORM) InitAllMileageByDegree() error {
	return db.Delete(&api.MileageByDegree{}).Error
}
