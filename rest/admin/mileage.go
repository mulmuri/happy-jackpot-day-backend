package admin

import (
	"backend/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitAllMileage(c *gin.Context) {
	if err := h.db.InitAllWeeklyMileage(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.db.InitAllMileageEarned(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.db.InitAllMileageByDegree(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//if err := h.db.InitAllUserMileage(); err != nil {
	//	c.String(http.StatusInternalServerError, err.Error())
	//	return
	//}

	c.String(http.StatusOK, "")
}

func (h *Handler) UpdateUserMileage(c *gin.Context) {
	var requests []api.Mileage
	var err error

	if err := c.ShouldBindJSON(&requests); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	for idx, request := range requests {
		if requests[idx].ID, err = h.db.GetUserID(request.Key); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	for _, request := range requests {
		if err := h.db.AddUserMileage(request.ID, request.Amount); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.String(http.StatusOK, "")
}
