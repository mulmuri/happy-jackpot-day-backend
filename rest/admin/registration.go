package admin

import (
	"backend/api"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) AcceptUserRegistration(c *gin.Context) {
	Key := c.Query("userid")
	fmt.Println(Key)
	user, err := h.db.GetUserByKey(Key)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if user.Status != api.UserStatusAwater {
		c.String(http.StatusBadRequest, errors.New("Not Awater").Error())
		return
	}

	if err := h.db.UpdateUserStatus(user.ID, api.UserStatusMember); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func (h *Handler) RejectUserRegistration(c *gin.Context) {
	Key := c.Query("userkey")
	user, err := h.db.GetUserByKey(Key)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.UpdateUserStatus(user.ID, api.UserStatusRejected); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func (h *Handler) HandleUserRegistration(c *gin.Context) {
	query := c.Query("accept")

	switch query {

	case "true":
		h.AcceptUserRegistration(c)
		return
	case "false":
		h.RejectUserRegistration(c)
		return
	default:
		err := errors.New("invalid parameters for accept")
		c.String(http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) GetAllRegisterRequests(c *gin.Context) {

	requests, err := h.db.GetAllRegisterRequests()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, requests)
}