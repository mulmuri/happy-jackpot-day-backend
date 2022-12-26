package admin

import (
	"backend/db"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	// userstatus
	GetAllRegisterRequests(c *gin.Context)
	HandleUserRegistration(c *gin.Context)
	AcceptUserRegistration(c *gin.Context)
	RejectUserRegistration(c *gin.Context)

	// mileage request
	GetAllMileageRequests(c *gin.Context)
	HandleMileageRequest(c *gin.Context)
	AcceptMileageRequest(c *gin.Context)
	RejectMileageRequest(c *gin.Context)
	UpdateUserMileage(c *gin.Context)

	// mileage propagation
	PropagateDailyMileage(c *gin.Context)
	PropagateDailyMileageByFile(c *gin.Context)
	InitAllMileage(c *gin.Context)

	// user
	GetAllUserList(c *gin.Context)
}

type Handler struct {
	db db.DBLayer
}

func NewHandler() HandlerInterface {
	return &Handler{
		db: db.DB,
	}
}