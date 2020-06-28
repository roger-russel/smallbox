.PHONY: dev
dev:
	@fresher -c .fresher.yaml

box-templates:
	@go run ./cmd/smallbox --force -f ./template/box/box.go.tpl
	@go run ./cmd/smallbox --force -f ./template/box/boxed.go.tpl

run-crud:
	@go run ./test/fixtures/crud/main.go
