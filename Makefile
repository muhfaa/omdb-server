.DEFAULT_GOAL=run

run:
	go run main.go

test:
	go test ./...

genproto:
	protoc \
		-I=./proto \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=./repository/grpcstub --go-grpc_opt=paths=source_relative \
		proto/*.proto proto/**/*.proto
