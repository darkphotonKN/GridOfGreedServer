# Lets "make" always run test targets
.PHONY: test 
	
build:
	 @go build -o game-server ./cmd/server/
	
run: build
	@./bin/starlight-cargo

test:
	@go test ./...

