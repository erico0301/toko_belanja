package controller

import (
	"net/http"
	"strconv"
	"toko_belanja/database"
	"toko_belanja/helper"
	"toko_belanja/model"

	"github.com/gin-gonic/gin"
)

func CategoryCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	Category := model.Category{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		c.ShouldBind(&Category)
	}
	Category.SoldProductAmount = 0
	err := db.Debug().Create(&Category).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                  Category.ID,
		"type":                Category.Type,
		"sold_product_amount": Category.SoldProductAmount,
		"created_at":          Category.SoldProductAmount,
	})
}

func CategoryViewAll(c *gin.Context) {
	db := database.GetDB()
	//contentType := helper.GetContentType(c)
	Category := []model.Category{}

	err := db.Find(&Category)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	for i := 0; i < len(Category); i++ {
		CategoryID := Category[i].ID

		Product := []model.Product{}
		err := db.Where("category_id=?", CategoryID).Find(&Product)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error,
			})
			return
		}
		Category[i].Product = Product
	}
	c.JSON(http.StatusOK, Category)
}

func CategoryUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	Category := model.Category{}
	CategoryID, _ := strconv.Atoi(c.Param("categoryId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		c.ShouldBind(&Category)
	}

	Category.ID = uint(CategoryID)

	db.Model(&Category).Where("id = ?", CategoryID).Update("type", Category.Type)

	db.First(&Category, CategoryID)

	c.JSON(http.StatusOK, gin.H{
		"id":                   Category.ID,
		"type":                 Category.Type,
		"sold_product_ammount": Category.SoldProductAmount,
		"updated_at":           Category.UpdatedAt,
	})
}

func CategoryDelete(c *gin.Context) {
	db := database.GetDB()
	CategoryID, _ := strconv.Atoi(c.Param("categoryId"))
	Category := model.Category{}
	db.Delete(&Category, CategoryID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}
