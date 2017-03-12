default: build

build:
	@go build .

clean:
	@go clean

.PHONY: default build clean
