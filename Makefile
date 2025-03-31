run: build
	@build/main

build: clean
	@mkdir -p build
	@go build -o ./build/main ./cmd/main.go

clean :
	@rm -rf build