package main

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/gookit/slog"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	staff_records "staff-ms/api/staff-records"
	"staff-ms/db"
	services "staff-ms/services"
)

var (
	HOST = os.Getenv("SVC_HOST")
	PORT = os.Getenv("SVC_PORT")
	TYPE = "tcp"
)

func registerSelf() {
	reqBody, _ := json.Marshal(map[string]string{
		"route": HOST + ":" + PORT,
		"svc":   "staff_svc",
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
	h := db.Init("staff-db", "staff_svc")

	lis, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	registerSelf()
	s := &services.Server{H: h}

	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)

	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
		))
	healthServer := health.NewServer()

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(staff_records.StaffRecords_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	staff_records.RegisterStaffRecordsServer(grpcServer, s)

	srvMetrics.InitializeMetrics(grpcServer)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go func() {
		if err := http.ListenAndServe(":40000", nil); err != nil {
			slog.Fatal(err)
		} else {
			slog.Info("HTTP server listening at 40000")
		}
	}()

	reflection.Register(grpcServer)

	log.Printf("services listening ar port %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
