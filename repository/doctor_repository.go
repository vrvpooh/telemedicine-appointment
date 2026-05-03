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

// ดึง doctor ตาม id
func GetDoctorByID(id string) (*model.Doctor, bool) {
	for _, d := range doctors {
		if d.ID == id {
			return &d, true
		}
	}
	return nil, false
}

// ดึง specialties ทั้งหมด
func GetAllSpecialties() []model.Specialty {
	return []model.Specialty{
		{Name: "Cardiology"},
		{Name: "Dermatology"},
		{Name: "Neurology"},
	}
}
