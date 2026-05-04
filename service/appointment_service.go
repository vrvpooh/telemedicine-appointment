package service

import (
	"errors"
	"strconv"
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

type AppointmentService struct {
	Repo     *repository.AppointmentRepository
	SlotRepo *repository.SlotRepository
}

func (s *AppointmentService) BookAppointment(patientID, slotID uint) (model.Appointment, error) {

	slot, err := s.SlotRepo.GetByID(slotID)
	if err != nil {
		return model.Appointment{}, err
	}
	if slot.IsBooked {
		return model.Appointment{}, errors.New("slot is already booked")
	}

	app := model.Appointment{
		PatientID: patientID,
		SlotID:    slotID,
		Status:    "confirmed",
		ZoomToken: "ZOOM-" + strconv.Itoa(int(slotID)),
	}

	if err := s.Repo.Create(&app); err != nil {
		return app, err
	}

	s.SlotRepo.UpdateIsBooked(slotID, true)
	return app, nil
}

func (s *AppointmentService) GetZoomToken(id uint) (string, error) {
	app, err := s.Repo.GetByID(id)
	return app.ZoomToken, err
}

func (s *AppointmentService) UpdateStatus(id uint, status string) error {
	app, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}
	app.Status = status
	return s.Repo.Update(app)
}
