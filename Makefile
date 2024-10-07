gen:
	protoc  --go_out=./internal --go-grpc_out=./internal internal/proto/*.proto;
	

# gen:
# 	protoc --proto_path=internal/proto internal/proto/*.proto --go_out=plugins=grpc:pb

generate_grpc:
	protoc \
    --go_out=.\
    --go_opt=paths=source_relative \
    --go-grpc_out=.  \
    --go-grpc_opt=paths=source_relative \
    internal/proto/**/*.proto

run_commander:
	go run ./cmd/spilothq/spilothq.go

# lint_proto:
    # protolint lint ./internal/proto/**/*.proto
