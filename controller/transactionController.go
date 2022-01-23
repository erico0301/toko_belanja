package controller

import (
	"net/http"
	"toko_belanja/database"
	"toko_belanja/helper"
	"toko_belanja/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TransactionCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	Transaction := model.Transaction{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Transaction)
	} else {
		c.ShouldBind(&Transaction)
	}

	//Cek Apakah product ada dan stok ada
	Product := model.Product{}
	err := db.First(&Product, Transaction.ProductID).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	if Transaction.Quantity > Product.Stock {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Product stock is not enough",
		})
		return
	}

	Transaction.TotalPrice = Transaction.Quantity * Product.Price

	//Cek user yg login
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	//Cek balance apakah cukup
	User := model.User{}
	err = db.First(&User, userID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	if Transaction.TotalPrice > int(User.Balance) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Saldo is not enough",
		})
		return
	}

	//Pengurangan stok product
	db.Model(&Product).Where("id = ?", Product.ID).Update("stock", Product.Stock-Transaction.Quantity)

	//Pengurangan saldo User
	db.Model(&User).Where("id = ?", User.ID).Update("balance", User.Balance-uint(Transaction.TotalPrice))

	//Create Transaction
	Transaction.UserID = userID

	err = db.Debug().Create(&Transaction).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Your have succesfully purchased the product",
		"transaction_bill": `{
			"total_price":Transaction.TotalPrice,
			"quantity":Transaction.Quantity,
			"Product_title":Product.Title,
		}`,
	})
}

func TransactionViewMyTransaction(c *gin.Context) {
	db := database.GetDB()
	//contentType := helper.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Transaction := []model.Transaction{}

	err := db.Where("user_id = ?", userID).Find(&Transaction)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	for i := 0; i < len(Transaction); i++ {
		Product := model.Product{}
		err := db.Where("id = ?", Transaction[i].ProductID).First(&Product)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error,
			})
			return
		}
		Transaction[i].Product = &Product
	}

	c.JSON(http.StatusOK, Transaction)
}

func TransactionViewAll(c *gin.Context) {
	db := database.GetDB()
	//contentType := helper.GetContentType(c)

	Transaction := []model.Transaction{}

	err := db.Find(&Transaction)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	for i := 0; i < len(Transaction); i++ {
		Product := model.Product{}
		err := db.Where("id = ?", Transaction[i].ProductID).First(&Product)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error,
			})
			return
		}
		Transaction[i].Product = &Product
	}

	c.JSON(http.StatusOK, Transaction)
}
