gen-sso-protos:
	protoc --proto_path=internal/proto/sso internal/proto/sso/sso.proto --go_out=internal/proto/sso/gen \
																		--go_opt=paths=source_relative \
																		--go-grpc_out=internal/proto/sso/gen \
																		--go-grpc_opt=paths=source_relative

run:
	go build -o ./bin/bin ./cmd/sso/main.go
	./bin/sso

migrate:
	go build -o ./bin/migrator ./cmd/migrator/main.go
	./bin/migrator --storage-path=./storage/sso.db --migrations-path=./migrations