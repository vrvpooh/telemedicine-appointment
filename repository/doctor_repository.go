package repository

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type DoctorRepository struct{}

// GET all doctors
func (r *DoctorRepository) GetAll() ([]model.Doctor, error) {
	rows, err := config.DB.Query("SELECT id, name, specialty FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []model.Doctor

	for rows.Next() {
		var d model.Doctor
		if err := rows.Scan(&d.ID, &d.Name, &d.Specialty); err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}

// GET doctor by ID
func (r *DoctorRepository) GetByID(id string) (*model.Doctor, error) {
	row := config.DB.QueryRow(
		"SELECT id, name, specialty FROM doctors WHERE id = ?", id,
	)

	var d model.Doctor
	if err := row.Scan(&d.ID, &d.Name, &d.Specialty); err != nil {
		return nil, err
	}

	return &d, nil
}

// GET specialties
func (r *DoctorRepository) GetSpecialties() ([]model.Specialty, error) {
	rows, err := config.DB.Query("SELECT DISTINCT specialty FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialties []model.Specialty

	for rows.Next() {
		var s model.Specialty
		if err := rows.Scan(&s.Name); err != nil {
			return nil, err
		}
		specialties = append(specialties, s)
	}

	return specialties, nil
}
