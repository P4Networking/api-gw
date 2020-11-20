.PHONY: all proto doc clean

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pisc-api-gw main.go

proto:
	@echo "\033[32m----- Compiling proto files -----\033[0m"
	mkdir -p proto/api
	protoc -I=./proto/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=proto/api --openapiv2_opt logtostderr=true proto/pisc.proto

doc:
	@echo "\033[32m----- You can view document in -----\033[0m"
	@echo "\033[32m----- http://localhost:6060/pkg/github.com/P4Networking/ -----\033[0m"
	@echo "\033[32m----- Press ctrl+c if you want to exit -----\033[0m"
	godoc -http=:6060

clean:
	@echo "\033[32m----- Clear all environment -----\033[0m"
	rm -r -f proto/api
	rm -f pisc-api-gw