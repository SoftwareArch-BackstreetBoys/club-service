tidy: 
	go mod tidy
	go mod vendor

install-tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

gen:
	$(GOPATH)/bin/oapi-codegen -generate fiber -package=api_gen -o ./http/gen/api_server_gen.go ./http/gen/api_spec.yaml
	$(GOPATH)/bin/oapi-codegen -generate types -package=api_gen -o ./http/gen/api_type_gen.go ./http/gen/api_spec.yaml

run:
	go run main.go

compose:
	docker-compose up