package db

import (
	"backend/api"
	"errors"
)

func (db *DBORM) GetMileage(userid uint64) (api.Mileage, error) {
	mileage := api.Mileage{ID: userid}
	return mileage, db.Find(&mileage).Error
}

func (db *DBORM) AddUserMileage(userid uint64, amountToAdd int) error {
	mileage := api.Mileage{ID: userid}
	if err := db.Find(&mileage).Error; err != nil {
		return err
	}

	newAmount := mileage.Amount + amountToAdd
	return db.Model(&mileage).Update("Amount", newAmount).Error
}

func (db *DBORM) SubtractUserMileage(userid uint64, amountToSubtract int) error {

	mileage := api.Mileage{ID: userid}
	if err := db.Model(mileage).Find(&mileage).Error; err != nil {
		return err
	}

	newAmount := mileage.Amount - amountToSubtract
	if newAmount < 0 {
		return errors.New("amount calculated is minus")
	}

	return db.Model(mileage).Update("Amount", newAmount).Error
}



func (db *DBORM) InitUserMileage(userid uint64) error {
	return db.Model(api.Mileage{ID: userid}).Update("amount", 0).Error
}

func (db *DBORM) InitAllUserMileage() error {
	return db.Model(api.Mileage{}).Where("1 = 1").Update("amount", 0).Error
}
