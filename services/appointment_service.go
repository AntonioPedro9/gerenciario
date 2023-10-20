package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type AppointmentService struct {
	appointmentRepository *repositories.AppointmentRepository
}

func NewAppointmentService(appointmentRepository *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{appointmentRepository}
}

func (as *AppointmentService) CreateAppointment(appointment *models.CreateAppointmentRequest) error {
	validAppointment := &models.Appointment{
		Date:     appointment.Date,
		BudgetID: appointment.BudgetID,
		UserID:   appointment.UserID,
	}

	return as.appointmentRepository.Create(validAppointment)
}

func (as *AppointmentService) ListAppointments(userID, tokenID uuid.UUID) ([]models.Appointment, error) {
	if userID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return as.appointmentRepository.List(userID)
}

func (as *AppointmentService) UpdateAppointment(appointment *models.UpdateAppointmentRequest, tokenID uuid.UUID) error {
	if appointment.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	existingAppointment, err := as.appointmentRepository.GetAppointmentById(appointment.ID)
	if err != nil {
		return err
	}
	if existingAppointment == nil {
		return utils.NotFoundError
	}

	validAppointment := &models.UpdateAppointmentRequest{
		ID:     appointment.ID,
		Date:   appointment.Date,
		UserID: appointment.UserID,
	}

	return as.appointmentRepository.UpdateAppointment(validAppointment)
}

func (as *AppointmentService) DeleteAppointment(appointmentID uint, authUserID uuid.UUID) error {
	existingAppointment, err := as.appointmentRepository.GetAppointmentById(appointmentID)
	if err != nil {
		return err
	}

	if existingAppointment.UserID != authUserID {
		return utils.UnauthorizedActionError
	}

	return as.appointmentRepository.DeleteAppointment(appointmentID)
}
