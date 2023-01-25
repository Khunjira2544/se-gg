// ตารางหลักการเบิกอุปกรณ์แพทย์  เหมือน ข้อมูลการรักษา
package controller

import (
	"net/http"

	"github.com/Khunjira2544/se-65-g08/entity"
	"github.com/gin-gonic/gin"
)

// POST Request มีทุกอันยกเว้น officer ไม่รู้ทำไม
func CreateRequest(c *gin.Context) {

	var request entity.Request
	var equipment entity.Equipment
	var location entity.Location

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id              //ของเราเป็น ค้นหา Equipment ด้วย id
	if tx := entity.DB().Where("id = ?", request.EquipmentID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id			//ของเราเป็น ค้นหา Location ด้วย id
	if tx := entity.DB().Where("id = ?", request.LocationID).First(&location); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 12: สร้าง 
	wv := entity.Request{
		Equipment: equipment, // โยงความสัมพันธ์กับ Entity 
		Location: location, // โยงความสัมพันธ์กับ Entity
	
		R_ID: request.R_ID,
		R_NAME: request.R_NAME,
		QUANTITY:    request.QUANTITY,
		TIME:         request.TIME,
		
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET
func GetRequest(c *gin.Context) {
	var request entity.Request
	id := c.Param("id")
	if err := entity.DB().Preload("Equipment").Preload("Location").Raw("SELECT * FROM requests WHERE id = ?", id).Scan(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": request})
}

// GET /
func ListRequest(c *gin.Context) {
	var request []entity.Request
	if err := entity.DB().Preload("Equipment").Preload("Location").Raw("SELECT * FROM requests").Find(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}

// DELETE /
func DeleteRequest(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM requests WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /
func UpdateRequest(c *gin.Context) {
	var request entity.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", request.ID).First(&request); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request not found"})
		return
	}

	if err := entity.DB().Save(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}
