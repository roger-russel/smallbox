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

.PHONY: test
test:
	@go test -v ./... -coverpkg="./box/...,./cmd/...,./internal/..." -cover -coverprofile=./coverage.txt -covermode=atomic -gcflags="all=-N -l"

.PHONY: coverage
coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
