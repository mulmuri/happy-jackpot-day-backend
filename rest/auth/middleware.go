package auth

import (
	"backend/api"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetAccessTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(ACCESS_SECRET), nil
}

func GetRefreshTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(REFRESH_SECRET), nil
}



func (h *Handler) SetAuthorityMiddleware() gin.HandlerFunc {
	return h.SetAuthority
}

func (h *Handler) SetAuthority(c *gin.Context) {
	accessTokenJWT, err := h.VerifyToken(c, "access-token");

	if err != nil {
		accessTokenJWT, err = h.RefreshToken(c);
		if err != nil {
			c.Set("status", api.UserStatusVisitor)
			c.Next()
			return;
		}
	}

	ts, err := h.ExtractTokenMetadata(accessTokenJWT)
	if err != nil {
		c.Set("status", api.UserStatusVisitor)
		c.Next();
		return
	}

	c.Set("userid", ts.UserId)
	c.Set("status", ts.UserStatus)
	c.Next();
}



func (h *Handler) AllowCorsMiddleware() gin.HandlerFunc {
	return h.AllowCors
}

func (h *Handler) AllowCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}




func (h *Handler) CheckAdminMiddleware() gin.HandlerFunc {
	return h.CheckAdmin
}

func (h *Handler) CheckUserMiddleware() gin.HandlerFunc {
	return h.CheckUser
}

func (h *Handler) CheckAdmin(c *gin.Context) {
	status := c.Keys["status"].(string)

	if status == api.UserStatusAdmin {
		c.Next()
		return
	}

	c.String(http.StatusUnauthorized, "Not an Admin")
	c.Abort()
}

func (h *Handler) CheckUser(c *gin.Context) {
	status := c.Keys["status"].(string)

    if status == api.UserStatusAdmin {
		c.Next()
		return
	}

	if status == api.UserStatusMember {
		c.Next()
		return
	}

	c.String(http.StatusUnauthorized, "Not an User")
	c.Abort()
}


