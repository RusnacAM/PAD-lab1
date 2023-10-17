package main

import (
	"bytes"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"schedule-ms/api/scheduler"
	"schedule-ms/db"
	"schedule-ms/services"
)

const (
	HOST = "schedule-ms"
	PORT = "3030"
	TYPE = "tcp"
)

func registerSelf() {
	reqBody, _ := json.Marshal(map[string]string{
		"route": HOST + ":" + PORT,
		"svc":   "scheduler_svc",
	})
	respBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post("http://service-discovery:5051/route", "application/json", respBody)

	if err != nil {
		log.Fatalf("An Error Ocurred %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func main() {
	h := db.Init()
	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	registerSelf()
	s := &services.Server{H: h}

	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(scheduler.Scheduler_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	scheduler.RegisterSchedulerServer(grpcServer, s)

	reflection.Register(grpcServer)
	log.Printf("services listening at port %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
