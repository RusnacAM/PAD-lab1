package test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	staff_records "staff-ms/api/staff-records"
	"testing"
)

type StaffServer struct {
	staff_records.UnimplementedStaffRecordsServer
}

func server(ctx context.Context) (staff_records.StaffRecordsClient, func()) {
	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	staff_records.RegisterStaffRecordsServer(baseServer, &StaffServer{})

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
					StaffID: "",
					Message: "Appointment created successfully",
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
				if tt.expected.out.Message != out.Message ||
					tt.expected.out.Error != out.Error {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}

}

//func TestGetStaffAvailability(t *testing.T) {
//	ctx := context.Background()
//
//	client, closer := server(ctx)
//	defer closer()
//
//	type expectation struct {
//		out *staff_records.GetAvailabilityResponse
//		err error
//	}
//
//	tests := map[string]struct {
//		in       *staff_records.GetStaffAvailability
//		expected expectation
//	}{
//		"Must_Success": {
//			in: &staff_records.CreateStaff{
//				Staff: staffRecord,
//			},
//			expected: expectation{
//				out: &staff_records.CreateResponse{
//					StaffID: "",
//					Message: "Appointment created successfully",
//					Error:   "",
//				},
//				err: nil,
//			},
//		},
//	}
//
//	for scenario, tt := range tests {
//		t.Run(scenario, func(t *testing.T) {
//			out, err := client.Create(ctx, tt.in)
//			if err != nil {
//				if tt.expected.err.Error() != err.Error() {
//					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
//				}
//			} else {
//				if tt.expected.out.Message != out.Message ||
//					tt.expected.out.Error != out.Error {
//					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
//				}
//			}
//
//		})
//	}
//
//}
