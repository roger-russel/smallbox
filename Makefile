.PHONY: dev
dev:
	@fresher -c .fresher.yaml

.PHONY: box-templates
box-templates:
	@go run ./cmd/smallbox --force -f ./template/box/box.go.tpl -n box
	@go run ./cmd/smallbox --force -f ./template/box/boxed.go.tpl -n boxed

.PHONY: run-crud
run-crud:
	@go run ./test/fixtures/crud/main.go
