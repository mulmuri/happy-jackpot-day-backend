package db

import (
	"backend/api"
	"errors"
)



func (db *DBORM) GetUserByID(userid uint64) (user api.User, err error) {
	result := db.Where("user_id = ?", userid).Find(&user)
	
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, errors.New("No User")
	}

	return user, nil
}

func (db *DBORM) GetUserByKey(userkey string) (user api.User, err error) {
	user.Key = userkey
	result := db.Where(user).Find(&user)
	
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, errors.New("No User")
	}

	return user, nil
}

func (db *DBORM) GetUserKey(userid uint64) (string, error) {
	user, err := db.GetUserByID(userid)
	return user.Key, err
}

func (db *DBORM) GetUserID(userkey string) (uint64, error) {
	user, err := db.GetUserByKey(userkey)
	return user.ID, err
}



func (db *DBORM) GetPersonalInfo(userid uint64) (userinfo api.PersonalInfo, err error) {
	result := db.Where("user_id = ?", userid).Find(&userinfo)
	
	if result.RowsAffected == 0 {
		return userinfo, errors.New("no user found")
	}

	if result.Error != nil {
		return userinfo, result.Error
	}

	userinfo.Key, err = db.GetUserKey(userid)
	if err != nil {
		return userinfo, err
	}

	userinfo.Recommender, err = db.GetRecommender(userid)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

func (db *DBORM) UpdatePersonalInfo(userinfo *api.PersonalInfo) error {
	result := db.Model(&api.PersonalInfo{ID: userinfo.ID}).Select("user_name", "email", "phone_number", "account_bank_name", "account_number").Updates(userinfo)

	if result.RowsAffected == 0 {
		return errors.New("no user found")
	}
	
	return result.Error
}
