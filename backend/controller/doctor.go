//แพทย์ เหมือน officer
package controller

import (
	"github.com/Khunjira2544/se-65-g08/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users   ของเราคือ Doctor
// GET /users
// List all users

func ListDoctors(c *gin.Context) {
	var doctors []entity.Doctor
	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctors).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doctors})
}

// GET /user/:id    ของเราคือ Doctor
// Get user by id
func GetDoctor(c *gin.Context) {
	var doctor entity.Doctor
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// PATCH /users    ของเราคือ Doctor
func UpdateDoctor(c *gin.Context) {
	var doctor entity.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", doctor.ID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	if err := entity.DB().Save(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// DELETE /users/:id
func DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM doctor WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
