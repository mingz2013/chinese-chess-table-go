help:
	@echo "Makefile help"

clean:
	rm chinese-chess-table-go -f


chinese-chess-table-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: chinese-chess-table-go
	docker build -t mingz2013/chinese-chess-table-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/chinese-chess-table-go


remove-container:
	docker rm chinese-chess-table-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name chinese-chess-table-go -it mingz2013/chinese-chess-table-go:latest


logs:
	docker logs chinese-chess-table-go


.PYONY: help, commit-docker, docker-image, chinese-chess-table-go, run, remove-container, logs

