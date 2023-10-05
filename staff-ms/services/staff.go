package services

import (
	"context"
	"log"
	staff_records "staff-ms/api/staff-records"
	"staff-ms/db"
)

type Server struct {
	H db.Handler
	staff_records.StaffRecordsServer
}

func (m *Server) Create(ctx context.Context, request *staff_records.CreateStaff) (*staff_records.CreateResponse, error) {
	log.Println("Create called")
	return &staff_records.CreateResponse{StaffID: "1", Message: "Appointment created successfully"}, nil
}

func (m *Server) Get(ctx context.Context, request *staff_records.GetStaffRecords) (*staff_records.GetResponse, error) {
	log.Println("Get called")
	return &staff_records.GetResponse{StaffRecords: []*staff_records.StaffRecord{
		{StaffID: "1", Name: "test name", JobTitle: "nurse", Department: "some dept", IsAvailable: true},
		{StaffID: "2", Name: "some name", JobTitle: "neurologist", Department: "some dept", IsAvailable: false},
		{StaffID: "3", Name: "first name", JobTitle: "cardiologist", Department: "some dept", IsAvailable: false},
		{StaffID: "4", Name: "last name", JobTitle: "nurse", Department: "some dept", IsAvailable: true},
	}}, nil
}

func (m *Server) GetAvailability(ctx context.Context, request *staff_records.GetStaffAvailability) (*staff_records.GetAvailabilityResponse, error) {
	log.Println("Get availability called")
	return &staff_records.GetAvailabilityResponse{IsAvailable: true}, nil
}

func (m *Server) Update(ctx context.Context, request *staff_records.UpdateStaff) (*staff_records.UpdateResponse, error) {
	log.Println("Update called")
	return &staff_records.UpdateResponse{StaffID: "1", Message: "Updated successfully"}, nil
}

func (m *Server) Delete(ctx context.Context, request *staff_records.DeleteStaff) (*staff_records.DeleteResponse, error) {
	log.Println("Delete called")
	return &staff_records.DeleteResponse{StaffID: "1", Message: "Deleted appointment"}, nil
}

func (m *Server) Check(_ context.Context, _ *staff_records.HealthCheckRequest) (*staff_records.HealthCheckResponse, error) {
	status := staff_records.HealthCheckResponse_SERVING

	return &staff_records.HealthCheckResponse{Status: status}, nil
}
