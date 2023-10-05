proto:
	python3 -m grpc_tools.protoc -I./svc --python_out=. --pyi_out=. --grpc_python_out=. ./svc/*/*.proto

serve:
	python3 main.py