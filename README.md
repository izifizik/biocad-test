# Biocad test task
Application that traks any changes with files on the server \
The app can also upload files to server and delete from it \
As for uploading files, server will not rewrite file with equal filename.\
\
Explain:
* If server don't have a file "filename.txt" - server save file like "filename.txt"
* If server have a file "filename.txt" - server save file like "filename(1).txt"
* If server have many file ("filename.txt", "filename(1).txt") - server save file like "filename(2).txt" and etc.

## Stack of technologies

* Go 1.17
* Mongodb
* Docker

## Handlers
### GET /
* Open a page where you can select and send a file to server
### GET /ws
* The handler where you can see what happen with the files (upload/delete) using websocket connection
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
