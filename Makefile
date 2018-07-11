.PHONY: build
build:
	middleman build
	cp _redirects build/_redirects
	mkdir -p functions
	go get ./...
	go build -o functions/analytics ./...
