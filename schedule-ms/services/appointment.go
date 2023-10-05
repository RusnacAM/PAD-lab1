package services

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"schedule-ms/api/scheduler"
	"schedule-ms/db"
	"schedule-ms/models"
)

type Server struct {
	H db.Handler
	scheduler.SchedulerServer
}

func (m *Server) CreateAppt(_ context.Context, request *scheduler.CreateAppointment) (*scheduler.CreateResponse, error) {
	log.Println("Create called")
	var appointment models.Appointment

	appointment.ApptID = uuid.New().String()
	appointment.PatientName = request.Appointment.PatientName
	appointment.DoctorName = request.Appointment.DoctorName
	appointment.ApptDateTime = request.Appointment.ApptDateTime
	appointment.ApptType = request.Appointment.ApptType

	if result := m.H.DB.Create(&appointment); result.Error != nil {
		return &scheduler.CreateResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	return &scheduler.CreateResponse{
		ApptID: appointment.ApptID,
		Status: http.StatusCreated,
	}, nil
}

func (m *Server) GetAppt(_ context.Context, _ *scheduler.GetAppointments) (*scheduler.GetResponse, error) {
	log.Println("Get called")
	appointments := []*scheduler.Appointment{}

	if result := m.H.DB.Find(&appointments); result.Error != nil {
		return &scheduler.GetResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	return &scheduler.GetResponse{Appointments: appointments, Status: http.StatusOK}, nil
}

func (m *Server) UpdateAppt(_ context.Context, request *scheduler.UpdateAppointment) (*scheduler.UpdateResponse, error) {
	log.Println("Update called")
	var appointment models.Appointment
	reqAppts := request.GetAppointment()

	if result := m.H.DB.Model(&appointment).Where("appt_id=?", reqAppts.ApptID).Updates(models.Appointment{
		PatientName:  reqAppts.PatientName,
		DoctorName:   reqAppts.DoctorName,
		ApptDateTime: reqAppts.ApptDateTime,
		ApptType:     reqAppts.ApptType,
	}); result.Error != nil {
		return &scheduler.UpdateResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	return &scheduler.UpdateResponse{Appointment: &scheduler.Appointment{
		ApptID:       appointment.ApptID,
		DoctorName:   appointment.DoctorName,
		PatientName:  appointment.PatientName,
		ApptDateTime: appointment.ApptDateTime,
		ApptType:     appointment.ApptType,
	}, Status: 0, Error: ""}, nil
}

func (m *Server) DeleteAppt(_ context.Context, request *scheduler.DeleteAppointment) (*scheduler.DeleteResponse, error) {
	log.Println("Delete called")
	var appointment models.Appointment
	reqID := request.GetApptID()

	if result := m.H.DB.Where("appt_id=?", reqID).Delete(&appointment); result.Error != nil {
		return &scheduler.DeleteResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	return &scheduler.DeleteResponse{ApptID: reqID, Status: http.StatusOK}, nil
}

func (m *Server) Check(_ context.Context, _ *scheduler.HealthCheckRequest) (*scheduler.HealthCheckResponse, error) {
	status := scheduler.HealthCheckResponse_SERVING

	return &scheduler.HealthCheckResponse{Status: status}, nil
}
