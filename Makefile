BINARY_NAME=cli-racing

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin .
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux .
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows .

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download