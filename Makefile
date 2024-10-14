gen:
	protoc  --go_out=. --go-grpc_out=. internal/proto/*.proto;
	
genapi:
	protoc  --go_out=app/spilothq/api --go-grpc_out=app/spilothq/api app/spilothq/api/v1/*.proto;

# gen:
# 	protoc --proto_path=internal/proto internal/proto/*.proto --go_out=plugins=grpc:pb


run_commander:
	go run ./app/spilothq/spilothq.go

devhq:
	air -c ./zarf/air/hq.toml
	
devadmin:
	air -c ./zarf/air/admin.toml

startui:
	cd ./app/spilotui && wails dev