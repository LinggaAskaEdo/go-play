.PHONY: build
build:
	@go mod tidy && \
	go mod vendor && \
	go build -o ./build/app ./src/cmd

.PHONY: run
run:
	@./build/app
	