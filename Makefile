RUN_COMMAND=docker-compose run app

go-test:
	${RUN_COMMAND} go test ./...
