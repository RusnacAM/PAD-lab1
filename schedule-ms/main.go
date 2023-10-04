package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"schedule-ms/api/scheduler"
	"schedule-ms/db"
	"schedule-ms/models"
)

const (
	HOST = "localhost"
	PORT = "3030"
	TYPE = "tcp"
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

	log.Println(appointment.ApptType, request.Appointment.GetApptType())

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
		ApptID:       reqAppts.ApptID,
		PatientName:  reqAppts.PatientName,
		DoctorName:   reqAppts.DoctorName,
		ApptDateTime: reqAppts.ApptDateTime,
		ApptType:     reqAppts.ApptType,
	}); result.Error != nil {
		return &scheduler.UpdateResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	log.Println(appointment)
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

func main() {
	h := db.Init()
	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := &Server{H: h}

	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(scheduler.Scheduler_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	scheduler.RegisterSchedulerServer(grpcServer, s)

	reflection.Register(grpcServer)
	log.Printf("services listening ar port %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
