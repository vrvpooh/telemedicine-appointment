package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

// business logic
func GetDoctors() []model.Doctor {
	doctors := repository.GetAllDoctors()

	// ตัวอย่าง logic (เพิ่มได้ภายหลัง)
	// เช่น filter, sort, validation

	return doctors
}

func GetDoctorByID(id string) (*model.Doctor, bool) {
	return repository.GetDoctorByID(id)
}

func GetSpecialties() []model.Specialty {
	return repository.GetAllSpecialties()
}
