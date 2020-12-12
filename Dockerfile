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
