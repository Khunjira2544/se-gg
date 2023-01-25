// ผู้ป่วย เหมือน teacher มี fkกับอักตาราง
package controller

import (
	"net/http"

	"github.com/Khunjira2544/se-65-g08/entity"
	"github.com/gin-gonic/gin"
)

// POST Patient
func CreatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": patient})
}

// GET Patient
func GetPatient(c *gin.Context) {
	var patient entity.Patient

	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET Patient
func ListPatients(c *gin.Context) {
	var patients []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patients").Find(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// //// อันนี้ลองดูดีๆ  สาขาวิชา  ข้างล่างยังไม่ได้ทำต่อ  ตรง owner_id ต้องได้แก้อีกทีว่ามันคืออะไร
/*func ListT_Facultys(c *gin.Context) {
	owner_id := c.Param("owner_id") //
	var teachers []entity.Teacher
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM teacher WHERE owner_id=?", owner_id).Find(&teachers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teachers})
}*/

// DELETE Patient
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patients WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH Patient
func UpdatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.ID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	if err := entity.DB().Save(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}
