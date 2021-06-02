# bookstore_users-api

1. Start up Mysql DB
```shell
docker-compose up -d mysql
```
2. Run locally users-api app
```shell
export MYSQL_USERNAME=root
export MYSQL_PASSWORD=root
export MYSQL_ADDRESS=localhost:3306
export MYSQL_SCHEMA=users_db
go run main.go
```
3. Run locally with docker-compose.
- Make sure the external network `shared-backend` is existing
```shell
docker network create shared-backend
```
- Run docker-compose
```
docker-compose up --build
```
