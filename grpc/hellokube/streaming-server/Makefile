build:
	GOOS=linux go build -o server main.go
	docker build -t grpc-server . && rm server
