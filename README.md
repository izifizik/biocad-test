#Biocad test task

## Stack of technologies

* Go 1.17
* Mongodb
* Docker

## Example of a .env file
```
PORT=port
MONGO_URL="mongo url"
```

## How to run app

`go run app/main.go`

## To run this in docker:
```
docker build -t "image name" .
docker run --name "container name" -p "local port":"container port" -d "image name"
docker exec -it containerID /bin/sh - for go incide container
```
