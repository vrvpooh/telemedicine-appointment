package handler

import (
	"net/http"
	"strconv"

	"telemedicine-api/model"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

var SlotSvc service.ISlotService

// SetupSlotHandler
func SetupSlotHandler(svc service.ISlotService) {
	SlotSvc = svc
}

// CreateSlot -> POST /api/doctor/:id/slots
func CreateSlot(c *gin.Context) {
	idParam := c.Param("id")
	doctorID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor id ต้องเป็นตัวเลข"})
		return
	}

	var slot model.Slot
	if err := c.ShouldBindJSON(&slot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := SlotSvc.CreateSlot(doctorID, slot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สร้าง slot สำเร็จ",
		"data":    created,
	})
}

// GetAvailableSlots -> GET /api/doctor/:id/slots
func GetAvailableSlots(c *gin.Context) {
	doctorID := c.Param("id")

	slots, err := SlotSvc.GetAvailableSlots(doctorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor_id": doctorID,
		"count":     len(slots),
		"slots":     slots,
	})
}

// DeleteSlot -> DELETE /api/slots/:id
func DeleteSlot(c *gin.Context) {
	id := c.Param("id")

	rows, err := SlotSvc.DeleteSlot(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบ slot id นี้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ลบ slot สำเร็จ",
		"id":      id,
	})
}
