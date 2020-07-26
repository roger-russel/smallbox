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

.PHONY: codeclimate # Must have CODECOV_TOKEN env set
codeclimate:
	@cc-test-reporter format-coverage -t gocov -p ${GOPATH} -d
	@cc-test-reporter upload-coverage

.PHONY: codecov # Must have CODECOV_TOKEN env set
codecov:
	@export CODECOV_TOKEN=${CODECOV_TOKEN_SMALLBOX}
	@curl -s https://codecov.io/bash > /tmp/codecov.sh && chmod +x /tmp/codecov.sh && bash /tmp/codecov.sh

.PHONY: release
release:
	@git tag -d v0.1
	@rm -fR dist
	@git tag v0.1
	@goreleaser release
