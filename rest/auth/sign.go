package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/api"
)

func (h *Handler) SignIn(c *gin.Context) {
	var requser api.User

	if err := c.ShouldBindJSON(&requser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.db.GetUserByKey(requser.Key)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	if err := CheckPassword(requser.PW, user.PW); err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	ats, rts, err := h.CreateToken(user.ID, user.Status)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.redis.CreateAuth(user.ID, ats, rts); err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}


	personal, err := h.db.GetPersonalInfo(user.ID);
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	user.PW = "";
	user.UserName = personal.UserName

	c.SetCookie("access-token",  ats.Token, accessTokenTime,  "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("refresh-token", rts.Token, refreshTokenTime, "/", "http://127.0.0.1:3000", false, true)

	c.JSON(http.StatusOK, user)
}

func (h *Handler) SignOut(c *gin.Context) {
	accessTokenJWT, err := h.VerifyToken(c, "access-token")
	if err == errors.New("token expired") {
		accessTokenJWT, err = h.RefreshToken(c)
	}

	metadata, err := h.ExtractTokenMetadata(accessTokenJWT)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.redis.DeleteTokens(metadata); err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("access-token",  "", 0, "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("refresh-token", "", 0, "/", "http://127.0.0.1:3000", false, true)

	c.String(http.StatusOK, "")
}
