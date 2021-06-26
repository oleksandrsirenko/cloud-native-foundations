# Go Hello World

This is a simple Go web application that prints a "Hello World" message.

## Run the application

This application listens on port `6111`

To run the application, use the following command:
```
go run main.go 
```

Access the application on: http://127.0.0.1:6111/

## Tasks


## Solution

1. Create a Dokerfile manualy or from the bash:
   `touch Dockerfile`
2. Open the Dockerfile and create layers due to the task #1 amd save it
3. Build an image (make sure that you are in the go-helloworld directory):
   `docker build -t go-helloworld .`

NOTE: If you get the permission denied responce:
   `Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post http://%2Fvar%2Frun%2Fdocker.sock/v1.24/build?buildargs=%7B%7D&cachefrom=%5B%5D&cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=Dockerfile&labels=%7B%7D&memory=0&memswap=0&networkmode=default&rm=1&shmsize=0&t=go-helloworld&target=&ulimits=null&version=1: dial unix /var/run/docker.sock: connect: permission denied`
   
   fix it:
   - `sudo usermod -aG docker ${USER}`
   - `su - ${USER}`
   
   check:
   - `docker` ps -a
   
   see details: [here](https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket)
   
   repeat step 3
   NOTE: if you have an error message like this:
   `go: go.mod file not found in current directory or any parent directory; see 'go help modules'`
   `The command '/bin/sh -c go build  -o helloworld' returned a non-zero code: 1`
   this mean you need to init go.mod file. We can handle this issue using one of the following approach:
   - add another one layer `RUN go mod init` before the `RUN go build  -o helloworld`

4. Run the Docker image: `docker run -d -p 6111:6111 go-helloworld`
5. Verify if the application is available and work properly on localhost 127.0.0.1:6111 (if you have problems accessing `https`, just change it to` http`)
6. Tag the app: `docker tag go-helloworld <YOUR DOCKERHUB REPO NAME>/go-helloworld:v1.0.0`
7. Push the app to the DockerHub repo (need account): `docker push <YOUR DOCKERHUB REPO NAME>/go-helloworld:v1.0.0`
   
Solved!