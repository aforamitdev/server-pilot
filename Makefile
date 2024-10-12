gen:
	protoc  --go_out=./internal --go-grpc_out=./internal internal/proto/*.proto;
	

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