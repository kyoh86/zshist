.PHONY: gen test lint install

VERSION := `git vertag get`
COMMIT  := `git rev-parse HEAD`

install: gen test lint
	go install -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" ./...

lint: test
	gometalinter ./...

test: gen
	go test -v --race ./...

gen:
	go generate ./...

