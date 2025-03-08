PROTO_DIR=api/todo-list
PROTO_FILE=$(PROTO_DIR)/todo-list.proto
ROOT_OUT_DIR=./

PROTOC=protoc

generate-pb:
	$(PROTOC) --go_out=$(ROOT_OUT_DIR) --go-grpc_out=$(ROOT_OUT_DIR) \
    		--go_opt=paths=source_relative \
    		--go-grpc_opt=paths=source_relative \
    		$(PROTO_FILE)

generate: generate-pb

clean:
	@rm -rf $(ROOT_OUT_DIR)/api/todo-list/*.pb.go

help:
	@echo "Доступные команды:"
	@echo "  make generate       - Сгенерировать protobuf и gRPC файлы"
	@echo "  make clean          - Удалить сгенерированные pb.go файлы"
