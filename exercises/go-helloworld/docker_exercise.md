# Go Hello World

*This is a simple Go web application that prints a "Hello World" message.*

## Run the application

This application listens on port `6111`

To run the application, use the following command:

```sh
go run main.go 
```

Access the application on: <http://127.0.0.1:6111/>

## Excercise 2: Docker for Application Packaging

Create the Docker image for the Go web application and push it to DockerHub, considering the following requirements:

### Task 1

Dockerfile:

- use the `golang:alpine` base image
- set the working directory to `/go/src/app`
- make sure to copy all the files from the current directory to the container working directory (e.g. `/go/src/app`)
- to build the application, use `go build -o helloworld` command, where `-o helloworld` will create the binary of the application with the name `helloworld`
- the application should be accessible on port `6111`
- and lastly, the command to start the container is to invoke the binary created earlier, which is `./helloworld`

### Task 2

Docker image:

- should have the name `go-helloworld`
- should have a valid tag, and a version with a major, minor, and patch included
- should be available in DockerHub, under your username e.g. `pixelpotato/go-helloworld`

### Task 3

Docker container:

- should be running on your local machine, by referencing the image from the DockerHub with a valid tag e.g. `pixelpotato/go-helloworld:v5.12.3`

:pushpin: **NOTE:** You will need to use `docker login` to login into Docker before pushing images to Docker Hub.
