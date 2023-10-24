package test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	staff_records "staff-ms/api/staff-records"
	db "staff-ms/db"
	"staff-ms/services"
	"testing"
)

func server(ctx context.Context) (staff_records.StaffRecordsClient, func()) {
	h := db.Init("0.0.0.0", "test_db")
	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)

	s := &services.Server{H: h}

	baseServer := grpc.NewServer()
	staff_records.RegisterStaffRecordsServer(baseServer, s)

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return listener.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := listener.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
		h.DB.Exec("TRUNCATE TABLE staff_records")
	}

	client := staff_records.NewStaffRecordsClient(conn)

	return client, closer

}
func TestCreateStaff(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *staff_records.CreateResponse
		err error
	}

	staffRecord := &staff_records.StaffRecord{
		Name:        "John Doe",
		JobTitle:    "Nurse",
		Department:  "Emergency",
		IsAvailable: true,
	}

	tests := map[string]struct {
		in       *staff_records.CreateStaff
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.CreateStaff{
				Staff: staffRecord,
			},
			expected: expectation{
				out: &staff_records.CreateResponse{
					Message: "Staff record created successfully",
					Error:   "",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Create(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Message != out.Message {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}

}

func TestGetStaff(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *staff_records.GetResponse
		err error
	}

	staffRecord := &staff_records.StaffRecord{
		Name:        "John Doe",
		JobTitle:    "Nurse",
		Department:  "Emergency",
		IsAvailable: true,
	}

	resp, err := client.Create(ctx, &staff_records.CreateStaff{Staff: staffRecord})
	if err != nil {
		panic(err)
	}

	tests := map[string]struct {
		in       *staff_records.GetStaffRecords
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.GetStaffRecords{},
			expected: expectation{
				out: &staff_records.GetResponse{
					StaffRecords: []*staff_records.StaffRecord{
						{
							StaffID:     resp.StaffID,
							Name:        staffRecord.Name,
							JobTitle:    staffRecord.JobTitle,
							Department:  staffRecord.Department,
							IsAvailable: staffRecord.IsAvailable,
						},
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Get(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.StaffRecords[0].StaffID != out.StaffRecords[0].StaffID ||
					tt.expected.out.StaffRecords[0].Name != out.StaffRecords[0].Name ||
					tt.expected.out.StaffRecords[0].JobTitle != out.StaffRecords[0].JobTitle ||
					tt.expected.out.StaffRecords[0].Department != out.StaffRecords[0].Department ||
					tt.expected.out.StaffRecords[0].IsAvailable != out.StaffRecords[0].IsAvailable {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}

}

func TestGetStaffAvailability(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *staff_records.GetAvailabilityResponse
		err error
	}

	staffRecord := &staff_records.StaffRecord{
		Name:        "John Doe",
		JobTitle:    "Nurse",
		Department:  "Emergency",
		IsAvailable: true,
	}

	resp, err := client.Create(ctx, &staff_records.CreateStaff{Staff: staffRecord})
	if err != nil {
		panic(err)
	}

	tests := map[string]struct {
		in       *staff_records.GetStaffAvailability
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.GetStaffAvailability{
				StaffID: resp.StaffID,
			},
			expected: expectation{
				out: &staff_records.GetAvailabilityResponse{
					StaffID:     resp.StaffID,
					IsAvailable: staffRecord.IsAvailable,
					Error:       "",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetAvailability(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.StaffID != out.StaffID ||
					tt.expected.out.IsAvailable != out.IsAvailable {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}

}

func TestUpdateStaff(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *staff_records.UpdateResponse
		err error
	}

	staffRecord := &staff_records.StaffRecord{
		Name:        "John Doe",
		JobTitle:    "Nurse",
		Department:  "Emergency",
		IsAvailable: true,
	}

	resp, err := client.Create(ctx, &staff_records.CreateStaff{Staff: staffRecord})
	if err != nil {
		panic(err)
	}

	updateRecord := &staff_records.StaffRecord{
		StaffID:     resp.StaffID,
		Name:        "Jane Doe",
		JobTitle:    "Radiologist",
		Department:  "Radiology",
		IsAvailable: false,
	}

	tests := map[string]struct {
		in       *staff_records.UpdateStaff
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.UpdateStaff{
				StaffRecord: updateRecord,
			},
			expected: expectation{
				out: &staff_records.UpdateResponse{
					StaffID: resp.StaffID,
					Message: "staff record successfully updated",
					Error:   "",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Update(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.StaffID != out.StaffID ||
					tt.expected.out.Message != out.Message {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}

}

func TestDeleteStaff(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *staff_records.DeleteResponse
		err error
	}

	staffRecord := &staff_records.StaffRecord{
		Name:        "John Doe",
		JobTitle:    "Nurse",
		Department:  "Emergency",
		IsAvailable: true,
	}

	resp, err := client.Create(ctx, &staff_records.CreateStaff{Staff: staffRecord})
	if err != nil {
		panic(err)
	}

	tests := map[string]struct {
		in       *staff_records.DeleteStaff
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.DeleteStaff{
				StaffID: resp.StaffID,
			},
			expected: expectation{
				out: &staff_records.DeleteResponse{
					StaffID: resp.StaffID,
					Message: "record successfully deleted",
					Error:   "",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Delete(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.StaffID != out.StaffID ||
					tt.expected.out.Message != out.Message {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}

}
