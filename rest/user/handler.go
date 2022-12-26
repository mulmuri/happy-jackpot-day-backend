package user

import (
	"backend/db"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {

	// registeration
	CheckIdValid(c *gin.Context)
	Register(c *gin.Context)
	Resign(c *gin.Context)

	// personalinfos
	GetUser(c *gin.Context)
	GetPersonalInfo(c *gin.Context)
	ChangePersonalInfo(c *gin.Context)

	// mileage
	GetMileage(c *gin.Context)
	GetWeeklyMileage(c *gin.Context)
	RequestMileageWithdrawal(c *gin.Context)
	GetMileageWithdrawalStatus(c *gin.Context)
	SaveMileage(c *gin.Context)

	// friend
	GetAllFriendsDailyMileage(c *gin.Context)
	GetAllMileageByDegree(c *gin.Context)
}

type Handler struct {
	db db.DBLayer
}

func NewHandler() HandlerInterface {
	return &Handler{
		db: db.DB,
	}
}
