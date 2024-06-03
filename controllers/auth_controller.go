package controllers

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/fernode/goBookStore/config"
	"github.com/fernode/goBookStore/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var identityKey = "id"

func getAuthMiddleware(db *gorm.DB) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "bookstore zone",
		Key:         []byte(config.EnvConfigs.SecretKey),
		Timeout:     24 * time.Hour,
		MaxRefresh:  24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Model: gorm.Model{ID: uint(claims[identityKey].(float64))},
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var user models.User
			if err := db.Where("username = ?", loginVals.Username).First(&user).Error; err != nil {
				return "", jwt.ErrFailedAuthentication
			}
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password)); err != nil {
				return "", jwt.ErrFailedAuthentication
			}
			return &user, nil
		},
	})
}

func Login(c *gin.Context) {
	authMiddleware, _ := getAuthMiddleware(c.MustGet("db").(*gorm.DB))
	authMiddleware.LoginHandler(c)
}
