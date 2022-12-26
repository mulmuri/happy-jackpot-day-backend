package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/api"
)


func (h *Handler) CheckIdValid(c *gin.Context) {

	var personalInfo api.PersonalInfo

	if err := c.ShouldBindJSON(&personalInfo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := personalInfo.CheckDataValidity(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.CheckUserOverlab(personalInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"flag": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"flag": true,
	})
}


func (h *Handler) Register(c *gin.Context) {

	var personalInfo api.PersonalInfo

	if err := c.ShouldBindJSON(&personalInfo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := personalInfo.CheckDataValidity(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.CheckUserOverlab(personalInfo); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	user := api.User{
		Key: personalInfo.Key,
		PW: personalInfo.PW,
		Status: api.UserStatusAwater,
	}

	if (c.Keys["status"].(string) == api.UserStatusAdmin) {
		user.Status = api.UserStatusMember
	}

	relation := api.Relation{
		RecommenderID: c.Keys["userid"].(uint64),
	}

	var err error
	if personalInfo.Recommender, err = h.db.GetUserKey(relation.RecommenderID); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.AddUser(&user, &personalInfo, &relation); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}



func (h *Handler) Resign(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	if err := h.db.UpdateUserStatus(userid, api.UserStatusSeceder); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
} 