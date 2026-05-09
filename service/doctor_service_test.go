package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDoctors(t *testing.T) {
	mockRepo := new(repository.MockDoctorRepository)

	mockData := []model.Doctor{
		{ID: "1", Name: "Dr. Smith"},
	}

	mockRepo.On("GetAll", "", "").Return(mockData, nil)

	service := DoctorService{Repo: mockRepo}

	result, err := service.GetDoctors("", "")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Dr. Smith", result[0].Name)
}

func TestGetDoctorByID(t *testing.T) {
	mockRepo := new(repository.MockDoctorRepository)

	mockDoctor := &model.Doctor{ID: "1", Name: "Dr. John"}

	mockRepo.On("GetByID", "1").Return(mockDoctor, nil)

	service := DoctorService{Repo: mockRepo}

	result, err := service.GetDoctorByID("1")

	assert.NoError(t, err)
	assert.Equal(t, "Dr. John", result.Name)
}

func TestGetSpecialties(t *testing.T) {
	mockRepo := new(repository.MockDoctorRepository)

	mockData := []model.Specialty{
		{Name: "Cardiology"},
	}

	mockRepo.On("GetSpecialties").Return(mockData, nil)

	service := DoctorService{Repo: mockRepo}

	result, err := service.GetSpecialties()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetDoctors_Error(t *testing.T) {
	mockRepo := new(repository.MockDoctorRepository)

	mockRepo.On("GetAll", "", "").Return([]model.Doctor{}, assert.AnError)

	service := DoctorService{Repo: mockRepo}

	_, err := service.GetDoctors("", "")

	assert.Error(t, err)
}

func TestGetDoctorByID_Error(t *testing.T) {
	mockRepo := new(repository.MockDoctorRepository)

	mockRepo.On("GetByID", "99").Return(&model.Doctor{}, assert.AnError)

	service := DoctorService{Repo: mockRepo}

	_, err := service.GetDoctorByID("99")

	assert.Error(t, err)
}

func TestGetDoctors_WithFilter(t *testing.T) {
	mockRepo := new(repository.MockDoctorRepository)

	mockRepo.On("GetAll", "smith", "Cardiology").Return([]model.Doctor{}, nil)

	service := DoctorService{Repo: mockRepo}

	_, err := service.GetDoctors("smith", "Cardiology")

	assert.NoError(t, err)
}
