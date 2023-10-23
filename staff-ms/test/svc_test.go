package test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	staff_records "staff-ms/api/staff-records"
	"staff-ms/db"
	"staff-ms/services"
	"testing"
)

type StaffServer struct {
	staff_records.UnimplementedStaffRecordsServer
}

func server(ctx context.Context) (staff_records.StaffRecordsClient, func()) {
	h := db.Init()
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
		StaffID:     "123",
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

	tests := map[string]struct {
		in       *staff_records.GetStaffRecords
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.GetStaffRecords{},
			expected: expectation{
				out: &staff_records.GetResponse{
					StaffRecords: []*staff_records.StaffRecord{},
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
				if tt.expected.out.StaffRecords != nil {
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

	tests := map[string]struct {
		in       *staff_records.GetStaffAvailability
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.GetStaffAvailability{
				StaffID: "5d4091a1-ea35-4844-8e04-8085acadb60a",
			},
			expected: expectation{
				out: &staff_records.GetAvailabilityResponse{
					StaffID:     "5d4091a1-ea35-4844-8e04-8085acadb60a",
					IsAvailable: true,
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
		StaffID:     "e9539d9c-791f-4779-8b67-c19b4f17ce63",
		Name:        "Test Update John Doe",
		JobTitle:    "Nurse",
		Department:  "Emergency",
		IsAvailable: true,
	}

	tests := map[string]struct {
		in       *staff_records.UpdateStaff
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.UpdateStaff{
				StaffRecord: staffRecord,
			},
			expected: expectation{
				out: &staff_records.UpdateResponse{
					StaffID: "e9539d9c-791f-4779-8b67-c19b4f17ce63",
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

	tests := map[string]struct {
		in       *staff_records.DeleteStaff
		expected expectation
	}{
		"Must_Success": {
			in: &staff_records.DeleteStaff{
				StaffID: "88286efe-d413-478c-b500-ccd55505b493",
			},
			expected: expectation{
				out: &staff_records.DeleteResponse{
					StaffID: "88286efe-d413-478c-b500-ccd55505b493",
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
