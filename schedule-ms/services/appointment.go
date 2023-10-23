package services

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"schedule-ms/api/scheduler"
	"schedule-ms/api/staff"
	"schedule-ms/db"
	"schedule-ms/models"
	"time"
)

type Server struct {
	H db.Handler
	scheduler.SchedulerServer
}

func getAvailability(staffID string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.Dial("staff-ms-0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := staff_records.NewStaffRecordsClient(conn)

	resp, err := client.GetAvailability(ctx, &staff_records.GetStaffAvailability{StaffID: staffID})
	if err != nil {
		panic(err)
	}

	return resp.GetIsAvailable()
}

func (m *Server) CreateAppt(_ context.Context, request *scheduler.CreateAppointment) (*scheduler.CreateResponse, error) {
	log.Println("Create called")
	var appointment models.Appointment

	appointment.ApptID = uuid.New().String()
	appointment.PatientName = request.Appointment.PatientName
	appointment.StaffID = request.Appointment.StaffID
	appointment.ApptDateTime = request.Appointment.ApptDateTime
	appointment.ApptType = request.Appointment.ApptType

	availability := getAvailability(appointment.StaffID)

	if !availability {
		return &scheduler.CreateResponse{Message: "The doctor you requested is not available, choose another."}, nil
	}

	if result := m.H.DB.Create(&appointment); result.Error != nil {
		return &scheduler.CreateResponse{Message: "The appointment could not be created due to an unexpected error.", Error: result.Error.Error()}, nil
	}

	return &scheduler.CreateResponse{
		ApptID:  appointment.ApptID,
		Message: "The appointment was created successfully",
	}, nil
}

func (m *Server) GetAppt(_ context.Context, _ *scheduler.GetAppointments) (*scheduler.GetResponse, error) {
	log.Println("Get called")
	appointments := []*scheduler.Appointment{}

	if result := m.H.DB.Find(&appointments); result.Error != nil {
		return &scheduler.GetResponse{Error: result.Error.Error()}, nil
	}

	return &scheduler.GetResponse{Appointments: appointments}, nil
}

func (m *Server) UpdateAppt(_ context.Context, request *scheduler.UpdateAppointment) (*scheduler.UpdateResponse, error) {
	log.Println("Update called")
	var appointment models.Appointment
	reqAppts := request.GetAppointment()

	availability := getAvailability(reqAppts.StaffID)

	if !availability {
		return &scheduler.UpdateResponse{Message: "Appointment could not be updated due to staff availability, choose another staff."}, nil
	}

	if result := m.H.DB.Model(&appointment).Where("appt_id=?", reqAppts.ApptID).Updates(models.Appointment{
		PatientName:  reqAppts.PatientName,
		StaffID:      reqAppts.StaffID,
		ApptDateTime: reqAppts.ApptDateTime,
		ApptType:     reqAppts.ApptType,
	}); result.Error != nil {
		return &scheduler.UpdateResponse{Message: "Appointment could not be updated.", Error: result.Error.Error()}, nil
	}

	return &scheduler.UpdateResponse{Appointment: &scheduler.Appointment{
		ApptID:       appointment.ApptID,
		StaffID:      appointment.StaffID,
		PatientName:  appointment.PatientName,
		ApptDateTime: appointment.ApptDateTime,
		ApptType:     appointment.ApptType,
	}, Message: "Appointment updated successfully."}, nil
}

func (m *Server) DeleteAppt(_ context.Context, request *scheduler.DeleteAppointment) (*scheduler.DeleteResponse, error) {
	log.Println("Delete called")
	var appointment models.Appointment
	reqID := request.GetApptID()

	if result := m.H.DB.Where("appt_id=?", reqID).Delete(&appointment); result.Error != nil {
		return &scheduler.DeleteResponse{Message: "Appointment could not be deleted.", Error: result.Error.Error()}, nil
	}

	return &scheduler.DeleteResponse{ApptID: reqID, Message: "Appointment deleted successfully."}, nil
}

func (m *Server) Check(_ context.Context, _ *scheduler.HealthCheckRequest) (*scheduler.HealthCheckResponse, error) {
	status := scheduler.HealthCheckResponse_SERVING

	return &scheduler.HealthCheckResponse{Status: status}, nil
}
