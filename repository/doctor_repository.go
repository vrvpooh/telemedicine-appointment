package repository

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type DoctorRepositoryInterface interface {
	GetAll(name, specialty string) ([]model.Doctor, error)
	GetByID(id string) (*model.Doctor, error)
	GetSpecialties() ([]model.Specialty, error)
}

type DoctorRepository struct{}

// GET all doctors
func (r *DoctorRepository) GetAll(name, specialty string) ([]model.Doctor, error) {
	query := "SELECT id, name, specialty, education, experience, rating FROM doctors WHERE 1=1"
	args := []interface{}{}

	if name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if specialty != "" {
		query += " AND specialty = ?"
		args = append(args, specialty)
	}

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []model.Doctor

	for rows.Next() {
		var d model.Doctor
		err := rows.Scan(&d.ID, &d.Name, &d.Specialty, &d.Education, &d.Experience, &d.Rating)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	if doctors == nil {
		doctors = []model.Doctor{}
	}

	return doctors, nil
}

// GET doctor by ID
func (r *DoctorRepository) GetByID(id string) (*model.Doctor, error) {
	row := config.DB.QueryRow(
		`SELECT id, name, specialty, education, experience, rating 
		 FROM doctors WHERE id = ?`, id,
	)

	var d model.Doctor
	err := row.Scan(&d.ID, &d.Name, &d.Specialty, &d.Education, &d.Experience, &d.Rating)
	if err != nil {
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
