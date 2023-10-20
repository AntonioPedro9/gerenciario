package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{db}
}

func (ar *AppointmentRepository) Create(appointment *models.Appointment) error {
	return ar.db.Create(appointment).Error
}

func (ar *AppointmentRepository) List(userID uuid.UUID) ([]models.Appointment, error) {
	var appointments []models.Appointment

	if err := ar.db.Where("user_id = ?", userID).Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

func (ar *AppointmentRepository) GetAppointmentById(id uint) (*models.Appointment, error) {
	var appointment models.Appointment

	if err := ar.db.Where("id = ?", id).First(&appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &appointment, nil
}

func (ar *AppointmentRepository) UpdateAppointment(appointment *models.UpdateAppointmentRequest) error {
	return ar.db.Model(&models.Appointment{}).
		Where("id = ?", appointment.ID).
		Updates(
			models.Appointment{
				Date: appointment.Date,
			},
		).Error
}

func (ar *AppointmentRepository) DeleteAppointment(appointmentID uint) error {
	appointment := models.Appointment{ID: appointmentID}
	return ar.db.Delete(&appointment).Error
}
