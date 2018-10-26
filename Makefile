help:
	@echo "Makefile help"

chinese-chess-table-go:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: chinese-chess-table-go
	docker build -t mingz2013/chinese-chess-table-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/chinese-chess-table-go


run:
	docker run --net="host" -it mingz2013/chinese-chess-table-go


.PYONY: help, commit-docker, docker-image, chinese-chess-table-go, run

