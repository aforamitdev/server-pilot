generate_grpc:
	protoc \
	--go_out=.\
	--go_opt=paths=source_relative \
	--go-grpc_out=.  \
	--go-grpc_opt=paths=source_relative \
	internal/proto/system/system.proto

run_commander:
	go run ./app/spilothq/spilothq.go

devhq:
	air -c ./zarf/air/hq.toml
	
devadmin:
	air -c ./zarf/air/admin.toml
	