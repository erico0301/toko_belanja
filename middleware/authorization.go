package middleware

import (
	"net/http"
	"toko_belanja/database"
	"toko_belanja/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CategoryAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := model.User{}

		err := db.Debug().Where("id=?", userID).Take(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := model.User{}

		err := db.Debug().Where("id=?", userID).Take(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func TransactionAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := model.User{}

		err := db.Debug().Where("id=?", userID).Take(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
