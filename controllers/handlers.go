package controllers

import (
	"dzabrail/connect"
	"dzabrail/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DrinkPreview struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetAllDrinks(c *gin.Context) {
	var drinks []DrinkPreview
	if err := connect.DB.Model(&models.Drink{}).Select("id", "name", "price").Find(&drinks).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"drinks": drinks})
}

func DrinksInstock(c *gin.Context) {
	var drink []models.Drink
	if err := connect.DB.Where("is_stock = ?", true).Find(&drink).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, drink)
}

func GetByID(c *gin.Context) {
	var drink models.Drink
	if err := connect.DB.First(&drink, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, drink)
}

func AddDrink(c *gin.Context) {
	var drink models.Drink
	if err := c.ShouldBindBodyWithJSON(&drink); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := connect.DB.Create(&drink).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, drink)
}

func DeleteDrink(c *gin.Context) {

	var drink models.Drink
	if err := connect.DB.First(&drink, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := connect.DB.Delete(&drink).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": true})
}

func UpdateDrink(c *gin.Context) {
	var drink models.Drink
	if err := connect.DB.First(&drink, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var inputDrink models.Drink

	if err := c.ShouldBindBodyWithJSON(&inputDrink); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := connect.DB.Model(&drink).Updates(inputDrink).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"update": true})
}
