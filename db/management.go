package db

import (
	"backend/api"
)



func (db *DBORM) UpdateUserStatus(userid uint64, status string) error {
	return db.Table("user").Where("user_id = ?", userid).Update("status", status).Error
}

func(db *DBORM) GetAllUserList() (requests []api.User, err error) {
	return requests, db.Find(&requests).Error
}



func(db *DBORM) GetAllRegisterRequests() ([]api.PersonalInfo, error) {
	requests := make([]api.PersonalInfo, 0)
	userlist := make([]api.User, 0)

	if err := db.Model(&userlist).Where(&api.User{Status: api.UserStatusAwater}).Find(&userlist).Error; err != nil {
		return requests, err
	}

	for _, user := range userlist {

		request, err := db.GetPersonalInfo(user.ID)
		if err != nil {
			return requests, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (db *DBORM) GetAllMemberList() ([]api.PersonalInfo, error) {
	list := make([]api.PersonalInfo, 0)
	userlist := make([]api.User, 0)

	if err := db.Model(&userlist).Where(&api.User{Status: api.UserStatusMember}).Find(&userlist).Error; err != nil {
		return list, err
	}

	for _, user := range userlist {
		userinfo, err := db.GetPersonalInfo(user.ID)
		if err != nil {
			return list, err
		}
		list = append(list, userinfo)
	}
	return list, nil
}