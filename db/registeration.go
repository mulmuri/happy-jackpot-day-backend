package db

import (
	"backend/api"
	"errors"
)

func (db *DBORM) CheckKeyOverlap(userkey string) error {
	result := db.Where("id = ?", userkey).Find(&api.User{})

	if result.RowsAffected != 0 {
		return errors.New("ID already exists.")
	}

	return result.Error
}

func (db *DBORM) CheckEmailOverlap(email string) error {
	result := db.Where("email = ?", email).Find(&api.PersonalInfo{})

	if result.RowsAffected != 0 {
		return errors.New("Email already exists.")
	}

	return result.Error
}

func (db *DBORM) CheckPhoneNoOverlap(phone string) error {
	result := db.Where("phone_number = ?", phone).Find(&api.PersonalInfo{})

	if result.RowsAffected != 0 {
		return errors.New("PhoneNumber already exists.")
	}

	return result.Error
}

func (db *DBORM) CheckUserOverlab(user api.PersonalInfo) error {
	if err := db.CheckKeyOverlap(user.Key); err != nil {
		return err
	}

	if err := db.CheckEmailOverlap(user.Email); err != nil {
		return err
	}

	if err := db.CheckPhoneNoOverlap(user.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (db *DBORM) AddUserAuth(user *api.User) (uint64, error) {
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}

	return db.GetUserID(user.Key)
}

func (db *DBORM) AddPersonalInfo(userinfo *api.PersonalInfo) error {
	return db.Create(&userinfo).Error
}

func (db *DBORM) AddMileageTuple(userid uint64) error {
	if err := db.Create(&api.Mileage{ID: userid}).Error; err != nil {
		return err
	}

	if err := db.Create(&api.WeeklyMileage{ID: userid}).Error; err != nil {
		return err
	}

	return nil
}

func (db *DBORM) AddUser(user *api.User, userinfo *api.PersonalInfo, relation *api.Relation) error {
	userID, err := db.AddUserAuth(user)
	if err != nil {
		return err
	}

	userinfo.ID = userID
	if err := db.AddPersonalInfo(userinfo); err != nil {
		return err
	}

	relation.ID = userID
	if err := db.AddUserRelation(relation); err != nil {
		return err
	}

	if err := db.AddMileageTuple(userID); err != nil {
		return err
	}

	return nil
}

func (db *DBORM) AddUserRelation(relation *api.Relation) error {
	return db.Create(&relation).Error
}
