# not working would need dockercompose to make it work.
# it does not work because mysql and redis service are not working.
# https://github.com/yun-mu/golang-graphql-example/blob/master/docker-compose.yml

FROM golang:1.15-alpine

ENV GOPATH /go

WORKDIR $GOPATH/src/graphql-srv

COPY . .
# RUN mkdir /app
# ADD . /app
# WORKDIR /app

RUN go mod download
RUN go mod verify

# RUN make run
EXPOSE 8080
CMD ["go","run", "server.go"]