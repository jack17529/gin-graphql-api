.PHONY: gqlgen

gqlgen:
	rm ./graph/schema.resolvers.go
	gqlgen generate

dev:
	go run server.go