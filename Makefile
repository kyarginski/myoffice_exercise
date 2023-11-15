GO111MODULE := auto
export GO111MODULE

lint:
	golangci-lint run ./...

test:
	go test -count=1 -race ./...

build:
	go build -tags musl -ldflags="-w -extldflags '-static' -X 'main.Version=$(VERSION)'" -o parser myoffice/cmd/parser

