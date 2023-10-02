package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	staff_records "staff-ms/api/staff-records"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type staffServer struct {
	staff_records.StaffRecordsServer
}

func (m *staffServer) Create(ctx context.Context, request *staff_records.CreateStaff) (*staff_records.CreateResponse, error) {
	log.Println("Create called")
	return &staff_records.CreateResponse{StaffID: "1", Message: "Appointment created successfully"}, nil
}

func (m *staffServer) Get(ctx context.Context, request *staff_records.GetStaffRecords) (*staff_records.GetResponse, error) {
	log.Println("Get called")
	return &staff_records.GetResponse{StaffRecords: []*staff_records.StaffRecord{
		{Name: "test name", JobTitle: "nurse", Department: "some dept", IsAvailable: true},
		{Name: "some name", JobTitle: "neurologist", Department: "some dept", IsAvailable: false},
		{Name: "first name", JobTitle: "cardiologist", Department: "some dept", IsAvailable: false},
		{Name: "last name", JobTitle: "nurse", Department: "some dept", IsAvailable: true},
	}}, nil
}

func (m *staffServer) GetAvailability(ctx context.Context, request *staff_records.GetStaffAvailability) (*staff_records.GetAvailabilityResponse, error) {
	log.Println("Get availability called")
	return &staff_records.GetAvailabilityResponse{IsAvailable: true}, nil
}

func (m *staffServer) Update(ctx context.Context, request *staff_records.UpdateStaff) (*staff_records.UpdateResponse, error) {
	log.Println("Update called")
	return &staff_records.UpdateResponse{StaffID: "1", Message: "Updated successfully"}, nil
}

func (m *staffServer) Delete(ctx context.Context, request *staff_records.DeleteStaff) (*staff_records.DeleteResponse, error) {
	log.Println("Delete called")
	return &staff_records.DeleteResponse{StaffID: "1", Message: "Deleted appointment"}, nil
}

func main() {
	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	staffServer := &staffServer{}
	staff_records.RegisterStaffRecordsServer(s, staffServer)

	log.Printf("server listening ar port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
