package api

import (
	"errors"
)



type User struct {
	ID       uint64 `gorm:"column:user_id"  json:"-"`
	Key      string `gorm:"column:id"       json:"ID"`
	PW       string `gorm:"column:password" json:"PW"`
	Status   string `gorm:"column:status"   json:"status"`
	UserName string `gorm:"-"               json:"name"`
}

func (User) TableName() string {
	return "user"
}



var UserStatusVisitor  string = "Visitor"
var UserStatusBanned   string = "Banned"
var UserStatusRejected string = "Rejected"
var UserStatusSeceder  string = "Seceder"
var UserStatusAwater   string = "Awater"
var UserStatusMember   string = "Member"
var UserStatusGuest    string = "Guest"
var UserStatusAdmin    string = "Admin"



func (u *User) CheckKeyValidity() error {
	if !(4 <= len(u.Key) && len(u.Key) <= 16) {
		return errors.New("User key range error")
	}

	return nil
}

func (u *User) CheckPWValidity() error {
	if !(4 <= len(u.PW) && len(u.PW) <= 16) {
		return errors.New("User pass range error")
	}

	return nil
}

func (u *User) CheckDataValidity() error {
	if err := u.CheckKeyValidity(); err != nil {
		return err
	}

	if err := u.CheckPWValidity(); err != nil {
		return err
	}
	
	return nil
}




