# Biocad test task

## Stack of technologies

* Go 1.17
* Mongodb
* Docker

## Handlers
### GET /
* Open a page where you can select and send a file to server
### GET /ws
* The page where you can see what happen with the files
### POST /upload
* Upload file to server
### POST /delete
* Delete file from server

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
docker exec -it containerID /bin/sh - for go inside container
```
