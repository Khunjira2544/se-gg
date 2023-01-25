//อุปกรณ์ เหมือน สถานะ
package controller

import (
	"net/http"

	"github.com/Khunjira2544/se-65-g08/entity"
	"github.com/gin-gonic/gin"
)

// POST Equipment
func CreateEquipment(c *gin.Context) {
	var equipment entity.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": equipment})
}

// GET Equipment
func GetEquipment(c *gin.Context) {
	var equipment entity.Equipment
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipment})
}

// GET Equipment
func ListEquipments(c *gin.Context) {
	var equipments []entity.Status
	if err := entity.DB().Raw("SELECT * FROM equipment").Scan(&equipments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})
}

// DELETE Equipment
func DeleteEquipment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM equipment WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH Equipment
func UpdateEquipment(c *gin.Context) {
	var equipment entity.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", equipment.ID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	if err := entity.DB().Save(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipment})
}
