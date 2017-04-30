FROM golang:1.8-alpine
LABEL maintainer "UshioShugo<ushio.s@gmail.com>"

ENV APP_PATH=${GOPATH}/src/github.com/ushios/dmm-clowrer-go

COPY . ${APP_PATH}

WORKDIR ${APP_PATH}

RUN apk add --no-cache --virtual .goget \
	git && \
	go get ./... && \
	apk del .goget


CMD ["go", "test", "./..."]