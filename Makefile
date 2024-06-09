

generate_grpc:
	protoc \
    --go_out=.\
    --go_opt=paths=source_relative \
    --go-grpc_out=.  \
    --go-grpc_opt=paths=source_relative \
    internal/proto/**/*.proto

run_commander:
	go run ./cmd/spilothq/spilothq.go

lint_proto:
    protolint lint ./internal/proto/**/*.proto