.PHONY: build run test clean

build2:
	@go build -ldflags "-s -w" -o raytracer .

build:
	@go build -o raytracer .

run: build
	@./raytracer

test:
	@go test ./...

clean:
	@rm -r ./raytracer ./test.png
