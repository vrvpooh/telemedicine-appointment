package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

type DoctorService struct {
	Repo repository.DoctorRepositoryInterface
}

func (s *DoctorService) GetDoctors(name, specialty string) ([]model.Doctor, error) {
	return s.Repo.GetAll(name, specialty)
}

func (s *DoctorService) GetDoctorByID(id string) (*model.Doctor, error) {
	return s.Repo.GetByID(id)
}

func (s *DoctorService) GetSpecialties() ([]model.Specialty, error) {
	return s.Repo.GetSpecialties()
}
