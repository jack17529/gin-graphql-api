# graphql-srv  
•	Constructed a high performance GraphQL server using Gin framework that stores video database using MySQL.  
•	Crafted mutations and queries for the server and modeled a caching service using Redis.  
• The main advantage of using graphql is that is provides more control of the query to the clients as they can get what the need and nothing more.  

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
  createVideo(input:{title:"video 2", url:"https://you.tube.com/vid2", userId:"2"}){
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
        "id": "2"
      },
      "title": "video 2",
      "url": "https://you.tube.com/vid2"
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
        "id": "2UWJUZOBCL",
        "title": "video 2",
        "url": "https://you.tube.com/vid2",
        "author": {
          "id": "2",
          "name": "user 2"
        }
      },
      {
        "id": "ZTUNJ5O5IQ",
        "title": "video 1",
        "url": "https://you.tube.com/vid1",
        "author": {
          "id": "1",
          "name": "user 1"
        }
      }
    ]
  }
}
```

Request
```
query GetVideoByID{
  getVideoById(video_id:"2UWJUZOBCL"){
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
    "getVideoById": {
      "id": "2UWJUZOBCL",
      "title": "video 2",
      "url": "https://you.tube.com/vid2",
      "author": {
        "id": "2",
        "name": "user 2"
      }
    }
  }
}
```

## MySQL

https://github.com/go-sql-driver/mysql  

Used MySQL database to keep the data.
1. `systemctl start mysql.service`
2. `systemctl status mysql.service`
3. `sudo mysql`
4. `show databases;` to show dbs.  
5. `create database videoDB;` to create videoDB  
6. `use videoDB;` to use mysql database.  
7. `show tables;` to see all the tables in the database.  
8. `create table videos (id varchar(50) not null primary key, title varchar(50), url varchar(70) not null unique, author_id varchar(50) not null, author_name varchar(50));` to create a table named videos.  
9. Set password for user root, the default is blank. After setting a password for `root` you will access mysql using `sudo mysql -u root -p`. 
10. Several ways to change root password `https://stackoverflow.com/questions/17975120/access-denied-for-user-rootlocalhost-using-password-yes-no-privileges` or `https://stackoverflow.com/questions/21944936/error-1045-28000-access-denied-for-user-rootlocalhost-using-password-y` 
11. `describe videos;` to see the structure of the table.
12. `select * from videos;` to see the table.
13. `exit` to get out of the db.


## Redis

https://github.com/go-redis/redis  

1. Use `redis-server` to start the redis server.
2. On another terminal type `redis-cli` to play with cli.
3. `select <INDEX>` to select the database.
4. `set key1 value1`
5. `get key1`
6. `append key1 1`
7. `keys *` to see all the keys.
8. `set key3 value3 ex 10` key will expire after 10 seconds.
9. `exit` to leave the cli.
