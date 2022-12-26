package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) GetAllUserList(c *gin.Context) {
	users, err := h.db.GetAllMemberList();
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}