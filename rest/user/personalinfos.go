package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/api"
)

func (h *Handler) GetUser(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	user, err := h.db.GetUserByID(userid);
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}


func (h *Handler) GetPersonalInfo(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	userinfo, err := h.db.GetPersonalInfo(userid);
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userinfo)
}



func (h *Handler) ChangePersonalInfo(c *gin.Context) {
    var userinfo = api.PersonalInfo{
		ID: c.Keys["userid"].(uint64),
	}

	if err := c.ShouldBindJSON(&userinfo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := userinfo.CheckDataValidity(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.UpdatePersonalInfo(&userinfo); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}