.PHONY: build
build:
	middleman build
	cp _redirects build/_redirects
	cp _headers build/_headers
	mkdir -p functions
	go get ./...
	go build -o functions/analytics ./...

.PHONY: run
run:
	./functions/analytics
