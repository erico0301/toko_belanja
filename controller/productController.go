package controller

import (
	"net/http"
	"strconv"
	"toko_belanja/database"
	"toko_belanja/helper"
	"toko_belanja/model"

	"github.com/gin-gonic/gin"
)

func ProductRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	Product := model.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Category := model.Category{}
	err := db.Debug().Where("id = ?", Product.CategoryID).Take(&Category).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "data not found",
			"message": "data doesn't exist",
		})
		return
	}

	err = db.Debug().Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          Product.ID,
		"title":       Product.Title,
		"price":       Product.Price,
		"stock":       Product.Stock,
		"category_id": Product.CategoryID,
		"created_at":  Product.CreatedAt,
	})

}

func ProductViewAll(c *gin.Context) {
	db := database.GetDB()

	Products := []model.Product{}

	err := db.Find(&Products)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	GetProduct := []model.GetProduct{}
	for i := 0; i < len(Products); i++ {
		tempGetProduct := model.GetProduct{}
		tempGetProduct.ID = Products[i].ID
		tempGetProduct.Title = Products[i].Title
		tempGetProduct.Price = Products[i].Price
		tempGetProduct.Stock = Products[i].Stock
		tempGetProduct.CategoryID = Products[i].CategoryID
		tempGetProduct.CreatedAt = Products[i].CreatedAt

		GetProduct = append(GetProduct, tempGetProduct)
	}
	c.JSON(http.StatusOK, GetProduct)
}

func ProductUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	Product := model.Product{}
	ProductID, _ := strconv.Atoi(c.Param("productId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}
	Category := model.Category{}
	err := db.Debug().Where("id = ?", Product.CategoryID).Take(&Category).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "data not found",
			"message": "data doesn't exist",
		})
		return
	}

	newProduct := model.Product{}
	_ = db.Where("id=?", ProductID).First(&newProduct)

	newProduct.Title = Product.Title
	newProduct.Price = Product.Price
	newProduct.Stock = Product.Stock
	newProduct.CategoryID = Product.CategoryID

	c.JSON(http.StatusOK, gin.H{
		"id":          newProduct.ID,
		"title":       newProduct.Title,
		"price":       newProduct.Price,
		"stock":       newProduct.Stock,
		"category_id": newProduct.CategoryID,
		"created_at":  newProduct.CreatedAt,
		"updated_at":  newProduct.UpdatedAt,
	})
}

func ProductDelete(c *gin.Context) {
	db := database.GetDB()
	Product := model.Product{}
	productID, _ := strconv.Atoi(c.Param("productId"))
	db.Delete(&Product, productID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Product has been successfully deleted",
	})
}
