.PHONY: build

BINARY_NAME=my-framework

build:
	go mod tidy
	GOARCH=amd64 GOOS=darwin go build -o build/package/${BINARY_NAME}-darwin cmd/server/main.go
	GOARCH=amd64 GOOS=linux go build -o build/package/${BINARY_NAME}-linux cmd/server/main.go
	GOARCH=amd64 GOOS=windows go build -o build/package/${BINARY_NAME}-windows cmd/server/main.go
	go build -o build/package/${BINARY_NAME} cmd/server/main.go

run: build
	./build/package/${BINARY_NAME}

clean:
	go clean
	rm build/package/${BINARY_NAME}-darwin
	rm build/package/${BINARY_NAME}-linux
	rm build/package/${BINARY_NAME}-windows
	rm build/package/${BINARY_NAME}

lint:
	golangci-lint run