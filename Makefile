.PHONY: dev
dev:
	@fresher -c .fresher.yaml

#It should run 2 times, one time to generate the b64 string
#and another one to generate the files with the b64 file updated
#it is necessary only into this project because it is self using it
.PHONY: box-templates
box-templates:
	@go run ./cmd/smallbox --force -f ./template/box/box.go.tpl -n box
	@go run ./cmd/smallbox --force -f ./template/box/box.go.tpl -n box
	@go run ./cmd/smallbox --force -f ./template/box/boxed.go.tpl -n boxed
	@go run ./cmd/smallbox --force -f ./template/box/boxed.go.tpl -n boxed

.PHONY: run-crud
run-crud:
	@go run ./test/fixtures/crud/main.go

.PHONY: test
test:
	@rm -fr ./cmd/smallbox/box
	@go test -v ./... -coverpkg="./box/...,./cmd/...,./internal/..." -cover -coverprofile=./coverage.txt -covermode=atomic -gcflags="all=-N -l"

.PHONY: coverage
coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
