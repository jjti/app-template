IMAGE_NAME ?= sblast

dev-api:
	@npx nodemon --watch './api/**/*.go' --ext go --signal SIGTERM --exec 'go' run .

dev-ui:
	@cd ui && npm run dev

deploy: ui
	@fly deploy

deps:
	@npm install
	@cd ui && npm install
	@go mod tidy
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install github.com/envoyproxy/protoc-gen-validate

.PHONY: proto
proto:
	@rm -rf ./gen
	@buf generate
	@npx openapi -i ./gen/sblast.swagger.json -o ./gen/client
