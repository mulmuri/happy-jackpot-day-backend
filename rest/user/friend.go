package user

import (
	"backend/api"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) GetAllFriendsDailyMileage(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	friends, err := h.db.GetFriendsList(userid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	friendsInfo := make([]api.FriendsInfo, len(friends))

	for idx := range friendsInfo {

		personal, err := h.db.GetPersonalInfo(friends[idx].ID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		friendsInfo[idx].UserName = personal.UserName
		friendsInfo[idx].Key = personal.Key

		mileage, err := h.db.GetMileage(friends[idx].ID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		friendsInfo[idx].Amount = mileage.Amount
	}

	c.JSON(http.StatusOK, friendsInfo)
}


func (h *Handler) GetAllMileageByDegree(c *gin.Context) {
	userid := c.Keys["userid"].(uint64)

	mileages, err := h.db.GetAllMileageByDegree(userid);
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mileages)
}