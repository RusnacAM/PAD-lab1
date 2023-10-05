proto:
	protoc ./api/staff-records/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

build:
	go build main.go

serve:
	go run main.go