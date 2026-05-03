package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

type DoctorService struct {
	Repo *repository.DoctorRepository
}

func (s *DoctorService) GetDoctors() ([]model.Doctor, error) {
	return s.Repo.GetAll()
}

func (s *DoctorService) GetDoctorByID(id string) (*model.Doctor, error) {
	return s.Repo.GetByID(id)
}

func (s *DoctorService) GetSpecialties() ([]model.Specialty, error) {
	return s.Repo.GetSpecialties()
}
