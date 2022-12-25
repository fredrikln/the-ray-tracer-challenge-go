.PHONY: build run test clean

build:
	@go build -o raytracer .

run: build
	@./raytracer

test:
	@go test ./...

clean:
	@rm -r ./raytracer ./test.png
