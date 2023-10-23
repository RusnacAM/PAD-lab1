package services

import (
	"context"
	"github.com/google/uuid"
	"log"
	staff_records "staff-ms/api/staff-records"
	"staff-ms/db"
	"staff-ms/models"
)

type Server struct {
	H db.Handler
	staff_records.StaffRecordsServer
}

func (m *Server) Create(ctx context.Context, request *staff_records.CreateStaff) (*staff_records.CreateResponse, error) {
	log.Println("Create called")
	var staff models.StaffRecord

	staff.StaffID = uuid.New().String()
	staff.Name = request.Staff.Name
	staff.JobTitle = request.Staff.JobTitle
	staff.Department = request.Staff.Department
	staff.IsAvailable = request.Staff.IsAvailable

	if result := m.H.DB.Create(&staff); result.Error != nil {
		return &staff_records.CreateResponse{
			Message: "New record couldn't be created due to unexpected error",
			Error:   result.Error.Error(),
		}, nil
	}

	return &staff_records.CreateResponse{StaffID: staff.StaffID, Message: "Staff record created successfully"}, nil
}

func (m *Server) Get(ctx context.Context, request *staff_records.GetStaffRecords) (*staff_records.GetResponse, error) {
	log.Println("Get called")
	staff := []*staff_records.StaffRecord{}

	if result := m.H.DB.Find(&staff); result.Error != nil {
		return &staff_records.GetResponse{Error: result.Error.Error()}, nil
	}

	return &staff_records.GetResponse{StaffRecords: staff}, nil
}

func (m *Server) GetAvailability(ctx context.Context, request *staff_records.GetStaffAvailability) (*staff_records.GetAvailabilityResponse, error) {
	log.Println("Get availability called")
	var staff models.StaffRecord
	reqID := request.GetStaffID()

	if result := m.H.DB.Find(&staff, "staff_id=?", reqID); result.Error != nil {
		return &staff_records.GetAvailabilityResponse{Error: result.Error.Error()}, nil
	}

	return &staff_records.GetAvailabilityResponse{StaffID: staff.StaffID, IsAvailable: staff.IsAvailable}, nil
}

func (m *Server) Update(ctx context.Context, request *staff_records.UpdateStaff) (*staff_records.UpdateResponse, error) {
	log.Println("Update called")
	var staff models.StaffRecord
	reqStaff := request.GetStaffRecord()
	log.Println(reqStaff.IsAvailable)

	if result := m.H.DB.Model(&staff).Where("staff_id=?", reqStaff.StaffID).Select("*").Updates(models.StaffRecord{
		StaffID:     reqStaff.StaffID,
		Name:        reqStaff.Name,
		JobTitle:    reqStaff.JobTitle,
		Department:  reqStaff.Department,
		IsAvailable: reqStaff.IsAvailable,
	}); result.Error != nil {
		return &staff_records.UpdateResponse{
			Message: "Staff record could not be updated due to unexpected error",
			Error:   result.Error.Error(),
		}, nil
	}

	log.Println(staff.IsAvailable)
	return &staff_records.UpdateResponse{
		StaffID: staff.StaffID,
		Message: "staff record successfully updated",
	}, nil
}

func (m *Server) Delete(ctx context.Context, request *staff_records.DeleteStaff) (*staff_records.DeleteResponse, error) {
	log.Println("Delete called")
	var staff models.StaffRecord
	reqID := request.GetStaffID()

	if result := m.H.DB.Where("staff_id=?", reqID).Delete(&staff); result.Error != nil {
		return &staff_records.DeleteResponse{Error: result.Error.Error()}, nil
	}
	return &staff_records.DeleteResponse{
		StaffID: reqID,
		Message: "record successfully deleted",
	}, nil
}

func (m *Server) Check(_ context.Context, _ *staff_records.HealthCheckRequest) (*staff_records.HealthCheckResponse, error) {
	status := staff_records.HealthCheckResponse_SERVING

	return &staff_records.HealthCheckResponse{Status: status}, nil
}
