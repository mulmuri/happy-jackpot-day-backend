package user

import (
	"backend/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWeeklyMileage(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	mileage, err := h.db.GetWeeklyMileage(userid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mileage)
}

func (h *Handler) GetMileage(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	mileage, err := h.db.GetMileage(userid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mileage)
}

func (h *Handler) RequestMileageWithdrawal(c *gin.Context) {
	request := api.MileageRequest{
		ID: c.Keys["userid"].(uint64),
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	mileage, err := h.db.GetMileage(request.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if mileage.Amount < request.Amount {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.AddMileageRequest(&request); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.db.SubtractUserMileage(request.ID, request.Amount); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func (h *Handler) GetMileageWithdrawalStatus(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	request, err := h.db.GetAllMileageRequestsByID(userid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, request)
}

func (h *Handler) SaveMileage(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)
	weekday := c.Query("weekday")

	if err := api.WeekdayValid(weekday); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	amount, err := h.db.TakeDailyMileage(userid, weekday)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.db.AddUserMileage(userid, amount); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}
