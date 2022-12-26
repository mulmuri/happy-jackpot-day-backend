package db

import (
	"backend/api"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBLayer interface {

	// registeration
	CheckUserOverlab(user api.PersonalInfo) error
	AddUser(user *api.User, userinfo *api.PersonalInfo, relation *api.Relation) error

	// management
	UpdateUserStatus(userid uint64, status string) error
	GetAllRegisterRequests() ([]api.PersonalInfo, error)
	GetAllUserList() ([]api.User, error)
	GetAllMemberList() ([]api.PersonalInfo, error)

	// mileage request
	GetAllMileageRequests() ([]api.MileageRequest, error)
	GetAllMileageRequestsByID(userid uint64) ([]api.MileageRequest, error)
	AcceptMileageRequest(requsetid uint64) error
	RejectMileageRequest(requestid uint64) error
	AddMileageRequest(request *api.MileageRequest) error

	// user
	GetUserByID(userid uint64) (api.User, error)
	GetUserKey(userid uint64) (string, error)
	GetUserByKey(userid string) (api.User, error)
	GetUserID(key string) (uint64, error)

	// relation
	GetRecommenderID(userid uint64) (uint64, error)
	GetRecommender(userid uint64) (string, error)
	GetFriendsList(userid uint64) ([]api.Relation, error)

	// personal info
	GetPersonalInfo(userid uint64) (api.PersonalInfo, error)
	UpdatePersonalInfo(user *api.PersonalInfo) error

	// weekly milege
	GetWeeklyMileage(userid uint64) (api.WeeklyMileage, error)
	InitAllWeeklyMileage() error

	// daily mileage
	TakeDailyMileage(userid uint64, weekday string) (int, error)
	AddDailyMileage(userid uint64, weekday string, amount int) error

	// mileage
	GetMileage(userid uint64) (api.Mileage, error)
	AddUserMileage(userid uint64, amount int) error
	SubtractUserMileage(userid uint64, amount int) error
	InitUserMileage(userid uint64) error
	InitAllUserMileage() error

	// mileage by degree
	AddDailyMileageByDegree(userid uint64, degree, amount int) error
	GetAllMileageByDegree(userid uint64) ([]api.MileageByDegree, error)
	InitAllMileageByDegree() error

	// mileage earned
	AddMileageEarned(userid uint64, amount int, weekday string) error
	InitAllMileageEarned() error
}



type DBORM struct {
	*gorm.DB
}

var DB *DBORM

func ConnectDB(address, dsn string) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:@(" + address + ")/management" + dsn,
	}))

	if err != nil {
		panic(err.Error())
	}

	DB = &DBORM{
		DB: db,
	}
}
