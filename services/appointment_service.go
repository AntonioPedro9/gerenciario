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

func (as *AppointmentService) UpdateAppointment(appointment *models.UpdateAppointmentRequest, tokenID uuid.UUID) (*models.Appointment, error) {
	if appointment.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	existingAppointment, err := as.appointmentRepository.GetAppointmentById(appointment.ID)
	if err != nil {
		return nil, err
	}
	if existingAppointment == nil {
		return nil, utils.NotFoundError
	}

	validAppointment := &models.UpdateAppointmentRequest{
		ID:     appointment.ID,
		Date:   appointment.Date,
		UserID: appointment.UserID,
	}

	updatedAppointment, err := as.appointmentRepository.Update(validAppointment)
	if err != nil {
		return nil, err
	}

	return updatedAppointment, nil
}

func (as *AppointmentService) DeleteAppointment(appointmentID uint, tokenID uuid.UUID) error {
	existingAppointment, err := as.appointmentRepository.GetAppointmentById(appointmentID)
	if err != nil {
		return err
	}

	if existingAppointment.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	return as.appointmentRepository.Delete(appointmentID)
}
