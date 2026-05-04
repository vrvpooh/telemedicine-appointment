package repository

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type SlotRepository struct{}

// Create slot
func (r *SlotRepository) Create(slot model.Slot) (model.Slot, error) {
	query := `INSERT INTO slots (doctor_id, start_time, end_time, is_booked)
	          VALUES (?, ?, ?, ?)`

	result, err := config.DB.Exec(query, slot.DoctorID, slot.StartTime, slot.EndTime, false)
	if err != nil {
		return slot, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return slot, err
	}

	slot.ID = id
	slot.IsBooked = false
	return slot, nil
}

// GetAvailableByDoctorID
func (r *SlotRepository) GetAvailableByDoctorID(doctorID string) ([]model.Slot, error) {
	query := `SELECT id, doctor_id, start_time, end_time, is_booked
	          FROM slots
	          WHERE doctor_id = ? AND is_booked = 0
	          ORDER BY start_time`

	rows, err := config.DB.Query(query, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	slots := []model.Slot{}
	for rows.Next() {
		var s model.Slot
		if err := rows.Scan(&s.ID, &s.DoctorID, &s.StartTime, &s.EndTime, &s.IsBooked); err != nil {
			return nil, err
		}
		slots = append(slots, s)
	}
	return slots, nil
}

// DeleteByID
func (r *SlotRepository) DeleteByID(id string) (int64, error) {
	result, err := config.DB.Exec("DELETE FROM slots WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
