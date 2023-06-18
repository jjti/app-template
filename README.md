# app-template

- platform: [fly.io](https://fly.io/)

- ui:

  - language: TypeScript
  - library: [React](https://react.dev/) / [Chakra](https://chakra-ui.com/)
  - bundler: [esbuild](https://esbuild.github.io/)
  - linter: [ESLint](https://eslint.org/docs/latest/)
  - formatter: [Prettier](https://prettier.io/)

- api:

  - data format: [proto3](https://protobuf.dev/programming-guides/proto3/) / [json](https://github.com/grpc-ecosystem/grpc-gateway)
  - validation: [protoc-gen-validate](https://github.com/bufbuild/protoc-gen-validate)
  - client gen: [swagger-typescript-api](https://www.npmjs.com/package/swagger-typescript-api)

- server:

  - language: Go
  - metrics: [prometheus](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus)

- db: [DynamoDB](https://docs.aws.amazon.com/dynamodb/index.html)
- infra: [Terraform](https://developer.hashicorp.com/terraform/docs)
