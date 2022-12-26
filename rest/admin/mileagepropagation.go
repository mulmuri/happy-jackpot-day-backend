package admin

import (
	"backend/api"
	"backend/util"
	"errors"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)



func (h *Handler) PropagateDailyMileage(c *gin.Context) {
	var requests []api.DailyMileageRequest

	if err := c.ShouldBindJSON(&requests); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var err error
	for idx, request := range requests {
		if requests[idx].ID, err = h.db.GetUserID(request.Key); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	for _, request := range requests {

		if err := h.db.AddMileageEarned(request.ID, request.Amount, request.Weekday); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if err := h.PropagateMileageToRecommender(request.ID, request.Amount, request.Weekday, 0, api.MileageRewardingDefaultLevel); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.String(http.StatusOK, "")
}



func (h *Handler) PropagateDailyMileageByFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if filepath.Ext(file.Filename) != ".xlsx" {
		err := errors.New("ext is not .xlsx")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	file.Filename = filepath.Join("./asset", time.Now().String() + ".xlsx")
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	updateInfos, err := util.ExtractDataFromExcel(file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	for idx, updateInfo := range updateInfos {
		_, err := h.db.GetUserByKey(updateInfo.Key)

		if err != nil {
			updateInfos[idx].State = api.DailyMileageReqStatusUserNotFound
		} else {
			updateInfos[idx].State = api.DailyMileageReqStatusValid
		}
	}

	c.JSON(http.StatusOK, updateInfos)
}


func (h *Handler) PropagateMileageToRecommender(user uint64, amount int, weekday string, depth, targetDepth int) error {
	if depth == targetDepth {
	    return nil
	}

	if err := h.db.AddDailyMileageByDegree(user, depth, amount); err != nil {
		return err
	}

	if err := h.db.AddDailyMileage(user, weekday, amount); err != nil {
		return err
	}

	recommender, err := h.db.GetRecommenderID(user)
	if err != nil {
		return err
	}

	if user == recommender {
		return nil
	}

	return h.PropagateMileageToRecommender(recommender, amount, weekday, depth+1, targetDepth)
}


