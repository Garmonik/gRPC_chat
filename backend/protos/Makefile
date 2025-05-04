.PHONY: auth chat gen

PROTO_DIR := proto
OUT_DIR := ./gen/go
COMMON_PROTOC_OPTS := \
	--go_out=$(OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(OUT_DIR) \
	--go-grpc_opt=paths=source_relative

gen: auth chat

auth:
	protoc -I $(PROTO_DIR) $(PROTO_DIR)/auth/*.proto $(COMMON_PROTOC_OPTS)

chat:
	protoc -I $(PROTO_DIR) $(PROTO_DIR)/chat/*.proto $(COMMON_PROTOC_OPTS)
