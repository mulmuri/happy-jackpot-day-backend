package admin

import (
	"backend/api"
	"net/http"
	"strconv"
	"errors"
	"github.com/gin-gonic/gin"
)


func (h *Handler) GetAllMileageRequests(c *gin.Context) {

	requests, err := h.db.GetAllMileageRequests()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, requests)
}

func (h *Handler) AcceptMileageRequest(c *gin.Context) {

	reqNo, err := strconv.ParseUint(c.Query("reqNo"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var request = api.MileageRequest{ReqId: reqNo}

	if err := h.db.AcceptMileageRequest(request.ReqId); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}


func (h *Handler) RejectMileageRequest(c *gin.Context) {

	reqNo, err := strconv.ParseUint(c.Query("reqNo"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var request = api.MileageRequest{ReqId: reqNo}

	if err := h.db.AcceptMileageRequest(request.ReqId); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.db.AddUserMileage(request.ID, request.Amount); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func (h *Handler) HandleMileageRequest(c *gin.Context) {
	query := c.Query("accept")

	switch query {
	case "true":
		h.AcceptMileageRequest(c)
		return
	case "false":
		h.RejectMileageRequest(c)
		return
	default:
		err := errors.New("invalid parameters for accept")
		c.String(http.StatusBadRequest, err.Error())
	}
}
