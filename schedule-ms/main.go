package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"schedule-ms/api/scheduler"
)

const (
	HOST = "localhost"
	PORT = "3030"
	TYPE = "tcp"
)

type scheduleServer struct {
	scheduler.SchedulerServer
}

func (m *scheduleServer) CreateAppt(ctx context.Context, request *scheduler.CreateAppointment) (*scheduler.CreateResponse, error) {
	log.Println("Create called")
	return &scheduler.CreateResponse{ApptID: "1", Status: "Appointment created successfully", Error: ""}, nil
}

func (m *scheduleServer) GetAppt(ctx context.Context, request *scheduler.GetAppointments) (*scheduler.GetResponse, error) {
	log.Println("Get called")

	return &scheduler.GetResponse{Appointments: []*scheduler.Appointment{
		{ApptID: "1", DoctorName: "sthsth", PatientName: "srgar", ApptDateTime: "argarg", Appointment: scheduler.Appointment_SURGICAL_INTERVENTION},
		{ApptID: "2", DoctorName: "sthsth", PatientName: "srgar", ApptDateTime: "argarg", Appointment: scheduler.Appointment_ROUTINE_CHECKUP},
		{ApptID: "3", DoctorName: "sthsth", PatientName: "srgar", ApptDateTime: "argarg", Appointment: scheduler.Appointment_SURGICAL_INTERVENTION},
		{ApptID: "4", DoctorName: "sthsth", PatientName: "srgar", ApptDateTime: "argarg", Appointment: scheduler.Appointment_SURGICAL_INTERVENTION},
	}}, nil
}

func (m *scheduleServer) UpdateAppt(ctx context.Context, request *scheduler.UpdateAppointment) (*scheduler.UpdateResponse, error) {
	log.Println("Update called")
	return &scheduler.UpdateResponse{Appointment: &scheduler.Appointment{
		ApptID: "1", DoctorName: "sthsth", PatientName: "srgar", ApptDateTime: "argarg", Appointment: scheduler.Appointment_SURGICAL_INTERVENTION,
	}, Status: "Updated successfully", Error: ""}, nil
}

func (m *scheduleServer) DeleteAppt(ctx context.Context, request *scheduler.DeleteAppointment) (*scheduler.DeleteResponse, error) {
	log.Println("Delete called")
	return &scheduler.DeleteResponse{ApptID: "1", Status: "Deleted appointment", Error: ""}, nil
}

func (m *scheduleServer) Check(ctx context.Context, request *scheduler.HealthCheckRequest) (*scheduler.HealthCheckResponse, error) {
	status := scheduler.HealthCheckResponse_SERVING

	return &scheduler.HealthCheckResponse{Status: status}, nil
}

func main() {
	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	apptServer := &scheduleServer{}
	healthServer := health.NewServer()

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(scheduler.Scheduler_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	scheduler.RegisterSchedulerServer(grpcServer, apptServer)

	reflection.Register(grpcServer)
	log.Printf("services listening ar port %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
