all: prepare build

test:
	@go test ./graph

prepare: 
	@echo "preparing..."

build: 
	@go build