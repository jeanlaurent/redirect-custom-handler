build:
	docker build . -t jeanlaurent/redirect-handler

run: build
	docker run -p 8080:8080 jeanlaurent/redirect-handler

dev:
	go run *.go
