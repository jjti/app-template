IMAGE_NAME ?= sblast

deploy: ui
	@fly deploy

.PHONY: ui
ui:
	@cd ui && npm run build

ui-dev:
	@cd ui && npm run dev

image:
	@docker build . -t $(IMAGE_NAME)

.PHONY: api
api:
	@go run .

deps:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install github.com/envoyproxy/protoc-gen-validate

.PHONY: proto
proto:
	@rm -rf ./gen ./ui/gen
	@buf generate
	@cd ui && npx openapi -i ../gen/sblast.swagger.json -o ./gen
