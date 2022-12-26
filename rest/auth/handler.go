package auth

import (
	"time"
	"backend/db"
	"backend/db/redis"
	"github.com/gin-gonic/gin"
)


var accessTokenTime  int = 1000 * 60 * 24
var refreshTokenTime int = 1000 * 60 * 24 * 14

var ACCESS_SECRET  = "access-secret"
var REFRESH_SECRET = "refresh-secret"

var AccessTokenExpiredAt  time.Duration = time.Hour * 24
var RefreshTokenExpiredAt time.Duration = time.Hour * 24 * 14



type HandlerInterface interface {

	// sign
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)

	// middleware
	SetAuthority(c *gin.Context)
	SetAuthorityMiddleware() gin.HandlerFunc
	AllowCors(c *gin.Context)
	AllowCorsMiddleware() gin.HandlerFunc
	CheckAdmin(c *gin.Context)
	CheckAdminMiddleware() gin.HandlerFunc
	CheckUser(c *gin.Context)
	CheckUserMiddleware() gin.HandlerFunc
}

type Handler struct {
	db db.DBLayer
	redis redis.RedisLayer
}

func NewHandler() HandlerInterface {
	return &Handler {
		db: db.DB,
		redis: redis.RDB,
	}
}
