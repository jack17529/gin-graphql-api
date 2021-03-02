# gin-graphql-api  
•	Constructed a high performance GraphQL API using Gin framework that stores video database using MySQL.  
•	Crafted mutations and queries for the server and modeled a caching service using Redis.  
The main advantage of using graphql is that is provides more control of the query to the clients as they can get what the need and nothing more.  

## NOTE
I have exposed the password for mysql and redis in this sample repository, in production they can be easily protected by using a `.env` file and supplying them.

## gqlgen
https://github.com/99designs/gqlgen  
Used `gqlgen` to generate the server code.

1. Use `go run github.com/99designs/gqlgen init` to initialize the graphql server.
2. Use `gqlgen generate` command to generate the code, in the `graph` directory.
3. `go run server.go` to run the server.

## Queries and Mutations

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

## Docker
1. Installation - https://www.ceos3c.com/hacking/install-docker-on-kali-linux/ depends on your os.
2. `sudo systemctl start docker`
3. `docker build -t <name of image you want to build>:<tag> .`
4. `sudo docker images` to see the current image.
5. `docker run -d -p 8080:8080 <name of image you want to build>:<tag>` pls run only in detach mode.
6. `dokcer ps -a` to see all teh containers.
7. `sudo docker logs -t <name of image you want to build>:<tag>` can be used to see the logs.
8. `sudo docker stop <name of image you want to build>:<tag>` to stop the container.

## docker-compose

1. `sudo docker-compose up`
2. `Ctrl+C` can be used to stop.
3. `sudo dokcer-compose down` to stop all the containers.
4. `sudo docker volume rm graphql-srv_mysql-vol` to remove the db volume container.

## mysql

1. `mysqldump -uroot -ppassword123 -d -B --events --routines --triggers videoDB > mysql_dump.sql` to dump the database.
2. `sudo docker cp mysql_dump.sql <CONTAINER_ID>:/mysql_dump.sql` to copy the dump file in the mysql docker container.
3. `sudo docker exec -it <CONTAINER_ID> /bin/bash` to get into the container.
4. `mysql -uroot -ppassword123 < mysql_dump.sql` to import the mysql database in the container.
5. `DESCRIBE videoDB.videos;` to check whether the import was successful. 
6. If you get Error 1130, then that means the host is not given the permissions from the database to connect.
7. To check use the `mysql` database in mysql container and query this `SELECT host,user FROM user;`
8. If the `root` user does not have it's host ip in there or the does not have `'%'` as host then it can't connect. 
9. To solve execute these commands
10. `CREATE USER 'root'@'%' IDENTIFIED BY '<yuour_password>';` to create a user `'root'@'%'` that can connect from any ip.
11. `GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';` to grant all previleges to `'root'@'%'` user.
12. Now if you check the output of the query, you will see this -
```
mysql> SELECT host,user FROM user;
+-----------+------------------+
| host      | user             |
+-----------+------------------+
| %         | root             |  <-- this :)
| 127.0.0.1 | root             |
| localhost | mysql.infoschema |
| localhost | mysql.session    |
| localhost | mysql.sys        |
| localhost | root             |
+-----------+------------------+
```
