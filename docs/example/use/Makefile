.PHONY: smallbox
smallbox:
	@smallbox -f ./assets/simple.txt -n simple

test:
	@go test -v -coverprofile=./coverage.txt -covermode=atomic -gcflags="all=-N -l"
