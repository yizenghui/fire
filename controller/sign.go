package controller

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	cpi "github.com/yizenghui/fire"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	OpenID string `json:"open_id"`
	Code   string `json:"code"`
	jwt.StandardClaims
}

// Sign 用户OPENID签名
func Sign(c echo.Context) error {
	code := c.QueryParam("code")
	ret, _ := cpi.GetOpenID(code)
	if code != "" && ret.OpenID != "" {

		// Set custom claims
		claims := &jwtCustomClaims{
			ret.OpenID,
			code,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// 获取签名里面的信息
func getOpenID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims.OpenID
}

// 获取用户信息
func getUser(openID string) (*cpi.Fans, error) {
	fans, err := cpi.GetFansByOpenID(openID)
	return fans, err
}

// Crypt 解密同步用户信息
func Crypt(c echo.Context) error {
	sessionKey := c.QueryParam("sk")
	encryptedData := c.QueryParam("ed")
	iv := c.QueryParam("iv")
	ret, _ := cpi.GetCryptData(sessionKey, encryptedData, iv)
	return c.JSON(http.StatusOK, ret)
}
