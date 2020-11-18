.PHONY: gqlgen

gqlgen:
	rm ./graph/schema.resolvers.go
	gqlgen generate

run:
	go run server.go

build:
	go build