package main

import (
	"context"
	"google.golang.org/grpc"
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

func (m *scheduleServer) Create(ctx context.Context, request *scheduler.CreateAppointment) (*scheduler.CreateResponse, error) {
	log.Println("Create called")
	return &scheduler.CreateResponse{ApptID: "1", Message: "Appointment created successfully"}, nil
}

func (m *scheduleServer) Get(ctx context.Context, request *scheduler.GetAppointments) (*scheduler.GetResponse, error) {
	log.Println("Get called")
	return &scheduler.GetResponse{Appointments: []*scheduler.Appointment{
		{DoctorID: "1", PatientID: "1", ApptDateTime: "argarg"},
		{DoctorID: "2", PatientID: "2", ApptDateTime: "vrbs"},
		{DoctorID: "3", PatientID: "3", ApptDateTime: "arbstbgarg"},
		{DoctorID: "4", PatientID: "4", ApptDateTime: "stnt"},
	}}, nil
}

func (m *scheduleServer) Update(ctx context.Context, request *scheduler.UpdateAppointment) (*scheduler.UpdateResponse, error) {
	log.Println("Update called")
	return &scheduler.UpdateResponse{ApptID: "1", Message: "Updated successfully"}, nil
}

func (m *scheduleServer) Delete(ctx context.Context, request *scheduler.DeleteAppointment) (*scheduler.DeleteResponse, error) {
	log.Println("Delete called")
	return &scheduler.DeleteResponse{ApptID: "1", Message: "Deleted appointment"}, nil
}

func main() {
	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	apptServer := &scheduleServer{}
	scheduler.RegisterSchedulerServer(s, apptServer)

	log.Printf("server listening ar port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
