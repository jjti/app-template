# Get creds from 1password: \
export \
AWS_REGION=us-east-1 \
AWS_ACCESS_KEY_ID=$(op read op://Private/AWS\ Personal/AWS_ACCESS_KEY_ID --account my.1password.com) \
AWS_SECRET_ACCESS_KEY=$(op read op://Private/AWS\ Personal/AWS_SECRET_ACCESS_KEY --account my.1password.com)
dev/server:
	@npx nodemon --watch '**/*.go' --ext go --signal SIGTERM --exec 'go' run .

dev/ui:
	@npm run dev

go/lint:
	@golangci-lint run api db

go/fix:
	@golangci-lint run api db --fix
	@go mod tidy

js/fix:
	@npm run fix

tf/fix:
	@terraform fmt -recursive .

deps:
	@npm install
	@go mod tidy
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

.PHONY: proto
proto:
	@cd pb && buf generate
	@npx swagger-typescript-api --path ./pb/service.swagger.json --output ./pb/client --name index.ts --api-class-name EchoService --silent

.PHONY: ui
ui:
	@npm run build

fly/deploy: ui
	@fly deploy

fly/dashboard:
	@fly dashboard

# Get creds from 1password: \
export \
AWS_REGION=us-east-1 \
AWS_ACCESS_KEY_ID=$(op read op://Private/AWS\ Personal/AWS_ACCESS_KEY_ID --account my.1password.com) \
AWS_SECRET_ACCESS_KEY=$(op read op://Private/AWS\ Personal/AWS_SECRET_ACCESS_KEY --account my.1password.com)
.PHONY: tf
tf:
	@cd tf && terraform apply
