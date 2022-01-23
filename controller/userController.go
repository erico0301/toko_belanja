package controller

import (
	"fmt"
	"net/http"
	"toko_belanja/database"
	"toko_belanja/helper"
	"toko_belanja/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password = User.Password
	err := db.Debug().Where("email=?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	User := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Role = "customer"
	User.Balance = 0

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"full_name":  User.FullName,
		"email":      User.Email,
		"password":   User.Password,
		"balance":    User.Balance,
		"created_at": User.CreatedAt,
	})
}

func UserTopUp(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	newUser := model.User{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	User := model.User{}
	_ = db.First(&User, userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&newUser)
	} else {
		c.ShouldBind(&newUser)
	}

	User.Balance = User.Balance + newUser.Balance
	fmt.Println("Balance : ", User.Balance)
	err := db.Model(&User).Where("id = ?", userID).Update("balance", &User.Balance).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	message := fmt.Sprint("Your balance has been successfully updated to Rp ", User.Balance)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
