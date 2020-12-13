.PHONY: all build openapi doc clean

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 5gc-api-gw main.go

openapi:
	@echo "\033[32m----- Compiling openapi files -----\033[0m"
	@mv openapi go
	java -jar ./tools/openapi-generator-cli.jar generate -i api/openapi.yaml -o . -g go-gin-server
	@mv go openapi
	@gsed -i 's/COPY go .\/go/COPY openapi .\/go/g' Dockerfile
	@gsed -i 's/.\/go/github.com\/P4Networking\/api-gw\/openapi/g' main.go
	@gsed -i 's/TraceDepth_2/TraceDepth2/g' openapi/model_trace_depth_2.go
	@gsed -i 's/REQUIRED/UpIntegrity_REQUIRED/g' openapi/model_up_integrity.go
	@gsed -i 's/PREFERRED/UpIntegrity_PREFERRED/g' openapi/model_up_integrity.go
	@gsed -i 's/NOT_NEEDED/UpIntegrity_NOT_NEEDED/g' openapi/model_up_integrity.go

doc:
	@echo "\033[32m----- You can view document in -----\033[0m"
	@echo "\033[32m----- http://localhost:6060/pkg/github.com/P4Networking/ -----\033[0m"
	@echo "\033[32m----- Press ctrl+c if you want to exit -----\033[0m"
	godoc -http=:6060

clean:
	@echo "\033[32m----- Clear all environment -----\033[0m"
	rm -r -f proto/api
	rm -f *-api-gw