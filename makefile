BINARY_NAME=local-diary
TARGET := ./server

# API=1
ifdef API 
	TARGET := ./API
endif

dep:
	go mod tidy

docs:
	godoc -http=:6060

swag:
	cd ${TARGET} && swag init -d ./,./controllers

storybook-dep:
	cd ./server/storybook && npm i

storybook:
	cd ./server/storybook && npm run storybook

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

tailwind:
	tailwindcss -i ./server/public/input.css -o ./server/public/dist/output.css

tailwind-watch:
	tailwindcss -i ./server/public/input.css -o ./server/public/dist/output.css --watch

webpack-dep:
	cd ./server/webpack && npm i

webpack:
	cd ./server/webpack && npm run build

webpack-watch:
	cd ./server/webpack && npm run watch

dev:
	go run ${TARGET} -swag -storybook ${ARGS}

#still need to add code to compile tailwind, and webpack
build:
	make tailwind
	make webpack
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ${TARGET} ${ARGS}
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${TARGET} ${ARGS}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ${TARGET} ${ARGS}

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows