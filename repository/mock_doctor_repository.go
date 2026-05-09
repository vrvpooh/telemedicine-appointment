package repository

import (
	"telemedicine-api/model"

	"github.com/stretchr/testify/mock"
)

type MockDoctorRepository struct {
	mock.Mock
}

func (m *MockDoctorRepository) GetAll(name, specialty string) ([]model.Doctor, error) {
	args := m.Called(name, specialty)
	return args.Get(0).([]model.Doctor), args.Error(1)
}

func (m *MockDoctorRepository) GetByID(id string) (*model.Doctor, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Doctor), args.Error(1)
}

func (m *MockDoctorRepository) GetSpecialties() ([]model.Specialty, error) {
	args := m.Called()
	return args.Get(0).([]model.Specialty), args.Error(1)
}
