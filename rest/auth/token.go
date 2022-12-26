package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"backend/api"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)



func (h *Handler) CreateToken(userid uint64, userstatus string) (atd api.TokenDetails, rtd api.TokenDetails, err error) {
	atd.Expires = time.Now().Add(AccessTokenExpiredAt).Unix()
	rtd.Expires = time.Now().Add(RefreshTokenExpiredAt).Unix()
	atd.Uuid = uuid.NewV4().String()
	rtd.Uuid = atd.Uuid + "++" + strconv.Itoa(int(userid))

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = atd.Uuid
	atClaims["user_id"] = userid
	atClaims["user_status"] = userstatus
	atClaims["exp"] = atd.Expires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	atd.Token, err = at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["refresh_uuid"] = rtd.Uuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = rtd.Expires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rtd.Token, err = rt.SignedString([]byte(REFRESH_SECRET))
	if err != nil {
		return
	}

	return atd, rtd, nil
}



func (h *Handler) VerifyToken(c *gin.Context, tokenType string) (*jwt.Token, error) {
	if !(tokenType == "access-token" || tokenType == "refresh-token") {
		return nil, errors.New("unexpected token type")
	}

	tokenString, err := c.Cookie(tokenType)
	if err != nil {
		return nil, err
	}

	var tokenParsed *jwt.Token
	switch tokenType {
	case "access-token":
		tokenParsed, err = jwt.Parse(tokenString, GetAccessTokenKey)
		break
	case "refresh-token":
		tokenParsed, err = jwt.Parse(tokenString, GetRefreshTokenKey)
		break
	default:
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if _, ok := tokenParsed.Claims.(jwt.Claims); !ok && !tokenParsed.Valid {
		return nil, errors.New("Token expired")
	}

	return tokenParsed, nil
}


func (h *Handler) RefreshToken(c *gin.Context) (*jwt.Token, error){
	tokenParsed, err := h.VerifyToken(c, "refresh-token")
	if err != nil {
		return nil, err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if (!(ok && tokenParsed.Valid)) {
		return nil, errors.New("can not get claims")
	}

	refreshUuid, _ := claims["refresh_uuid"].(string)

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	userStatus, _ := claims["user_status"].(string)

	if deleted, err := h.redis.DeleteAuth(refreshUuid); err != nil || deleted == 0 {
		return nil, err
	}

	ats, rts, err := h.CreateToken(userId, userStatus)
	if err != nil {
		return nil, err
	}
	
	if err := h.redis.CreateAuth(userId, ats, rts); err != nil {
		return nil, err
	}

	c.SetCookie("access-token",  ats.Token, accessTokenTime,  "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("refresh-token", rts.Token, refreshTokenTime, "/", "http://127.0.0.1:3000", false, true)
	
	token, err := h.VerifyToken(c, "access-token");
	if err != nil {
		return nil, err
	}

	return token, nil
}


func (h *Handler) ExtractTokenMetadata(token *jwt.Token) (*api.AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("")
	}

	accessUuid, ok := claims["access_uuid"].(string)
	if !ok {
		return nil, errors.New("")
	}

	UserStatus, ok := claims["user_status"].(string)
	if !ok {
		return nil, errors.New("")
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	return &api.AccessDetails{
		AccessUuid: accessUuid,
		UserId:     userId,
		UserStatus: UserStatus,
	}, nil
}




