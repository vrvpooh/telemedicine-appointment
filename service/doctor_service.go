package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

// business logic อยู่ตรงนี้
func GetDoctors() []model.Doctor {
	doctors := repository.GetAllDoctors()

	// ตัวอย่าง logic (เพิ่มได้ภายหลัง)
	// เช่น filter, sort, validation

	return doctors
}
