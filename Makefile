PROTO_PATH=./api/proto
PROTO=$(PROTO_PATH)/*.proto
PROTO_OUT=./internal/genproto
GATEWAY_OUT=$(PROTO_OUT)
GATEWAY_API=api/gateway/note.yaml
SWAGGER=./api/swagger

proto:
	protoc -I $(PROTO_PATH) \
		--go_out=$(PROTO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO)

grpc-gateway:
	protoc -I $(PROTO_PATH) \
		--grpc-gateway_out $(GATEWAY_OUT)\
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=$(GATEWAY_API) \
		$(PROTO)

swagger:
	protoc -I $(PROTO_PATH) \
		--openapiv2_out $(SWAGGER) \
		--openapiv2_opt grpc_api_configuration=$(GATEWAY_API) \
		$(PROTO)

tidy:
	go mod tidy

install-deps: tidy
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate: proto grpc-gateway swagger
