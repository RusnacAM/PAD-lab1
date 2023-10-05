package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	staff_records "staff-ms/api/staff-records"
	"staff-ms/db"
	services "staff-ms/services"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	h := db.Init()

	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := &services.Server{H: h}

	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(staff_records.StaffRecords_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	staff_records.RegisterStaffRecordsServer(grpcServer, s)

	reflection.Register(grpcServer)

	log.Printf("services listening ar port %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
