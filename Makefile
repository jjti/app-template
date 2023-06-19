# dev
dev:
	@npx nodemon --watch '**/*.go' --watch '**/*.ts' --watch '**/*.tsx' --ext go,ts,tsx --signal SIGTERM --exec 'npm run build && go run ./cmd'

# go
go/lint:
	@golangci-lint run --config ./.golangci.yaml

go/fix:
	@golangci-lint run --fix --config ./.golangci.yaml
	@go mod tidy

# js
js/fix:
	@npm run fix

# api
.PHONY: proto
proto:
	@cd pb && buf generate
	@npx swagger-typescript-api --path ./pb/service.swagger.json --output ./pb/client --name index.ts --api-class-name EchoService --silent

# fly
fly/deploy:
	@npm run build
	@fly deploy

fly/dashboard:
	@fly dashboard

# terraform
tf/apply:
	@cd tf && terraform apply

tf/fix:
	@terraform fmt -recursive .

# deps
deps:
	@go mod tidy
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
