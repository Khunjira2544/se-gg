//เทคนิคการแพทย์ เหมือน แพทย์
package controller

import (
	"github.com/Khunjira2544/se-65-g08/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users   ของเราคือ Medicine
// GET /users
// List all users

func ListMedicines(c *gin.Context) {
	var medicines []entity.Medicine
	if err := entity.DB().Raw("SELECT * FROM medicines").Scan(&medicines).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicines})
}

// GET /user/:id    ของเราคือ Medicine
// Get user by id
func GetMedicine(c *gin.Context) {
	var medicine entity.Medicine
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// PATCH /users    ของเราคือ Medicine
func UpdateMedicine(c *gin.Context) {
	var medicine entity.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medicine.ID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	if err := entity.DB().Save(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// DELETE /users/:id
func DeleteMedicine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
