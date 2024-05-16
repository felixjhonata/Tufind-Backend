package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetTutors() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tutors []models.Tutor
		database.DB.Find(&tutors)
		c.JSON(http.StatusOK, gin.H{"tutors": tutors})

	}
}

func GetTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tutor models.Tutor
		id := c.Param("id")
		if err := database.DB.First(&tutor, id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"tutor": tutor})
	}
}

func CreateTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tutor models.Tutor
		if err := c.ShouldBindJSON(&tutor); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		database.DB.Create(&tutor)
		c.JSON(http.StatusOK, gin.H{"tutor": tutor})
	}
}

func UpdateTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tutor models.Tutor
		id := c.Param("id")
		if err := c.ShouldBindJSON(&tutor); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if database.DB.Model(&tutor).Where("id = ?", id).Updates(&tutor).RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "can't update tutor"})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data update Success"})

	}
}

// func round(num float64) int {

// }

// func toFixed(num float64, precision int) float64 {

// }
