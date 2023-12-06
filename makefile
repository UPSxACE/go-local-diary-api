BINARY_NAME=local-diary
TARGET := ./server

# API=1
ifdef API 
	TARGET := ./API
endif

dep:
	go mod tidy

test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

watch:
	tailwindcss -i ./server/public/input.css -o ./server/public/dist/output.css --watch

dev:
	go run ${TARGET}

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ${TARGET}
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${TARGET}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ${TARGET}

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows