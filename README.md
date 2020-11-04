# gin-graphql-srv

Implements a graphql server which uses Gin http framework for serving the platform for adding videoes and getting the list of all the videos till yet. For Database I might use MySQl with Redis cache.  
The main advantage of using graphql is that is provides more control of the query to the clients as they can get what the need and nothing more.

## gqlgen

https://github.com/99designs/gqlgen  
Used `gqlgen` to generate the server code.

1. Use `go run github.com/99designs/gqlgen init` to initialize the graphql server.
2. Use `gqlgen generate` command to generate the code, in the `graph` directory.
3. `go run server.go` to run the server.

## Queries

Request
```
mutation createVideo{
  createVideo(input:{title:"video 1", url:"https://you.tube.com/vid1", userId:"1"}){
    author{
      id
    }
    title
    url
  }
}
```
As we did not ask for the id of the video and the name of the author, so we did not receive it below.

Response
```
{
  "data": {
    "createVideo": {
      "author": {
        "id": "1"
      },
      "title": "video 1",
      "url": "https://you.tube.com/vid1"
    }
  }
}
```

Request
```
query GetVideos{
  getVideos{
    id
    title
    url
    author{
      id
      name
    }
  }
}
```

Response

```
{
  "data": {
    "getVideos": [
      {
        "id": "T5577006791947779410",
        "title": "video 1",
        "url": "https://you.tube.com/vid1",
        "author": {
          "id": "1",
          "name": "user 1"
        }
      },
      {
        "id": "T8674665223082153551",
        "title": "video 2",
        "url": "https://you.tube.com/vid2",
        "author": {
          "id": "2",
          "name": "user 2"
        }
      }
    ]
  }
}
```

