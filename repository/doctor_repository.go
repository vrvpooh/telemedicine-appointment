package repository

import "telemedicine-api/model"

//mock data
var doctors = []model.Doctor{
	{ID: "1", Name: "Dr. Smith", Specialty: "Cardiology"},
	{ID: "2", Name: "Dr. John", Specialty: "Dermatology"},
	{ID: "3", Name: "Dr. Lee", Specialty: "Neurology"},
}

// ดึง doctor ทั้งหมด
func GetAllDoctors() []model.Doctor {
	return doctors
}
