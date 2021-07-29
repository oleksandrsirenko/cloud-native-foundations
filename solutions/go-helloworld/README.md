# Go Hello World

---

***This is a simple Go web application that prints a "Hello World" message.***

## Run the application

---

This application listens on port `6111`

To run the application, use the following command:
```
go run main.go 
```

Access the application on: http://127.0.0.1:6111/


## Excercise 1: Docker for Application Packaging

---

Create the Docker image for the Go web application and push it to DockerHub, considering the following requirements:

### Task #1
Dockerfile:

- use the `golang:alpine` base image
- set the working directory to `/go/src/app`
- make sure to copy all the files from the current directory to the container working directory (e.g. `/go/src/app`)
- to build the application, use `go build -o helloworld` command, where `-o helloworld` will create the binary of the application with the name `helloworld`
- the application should be accessible on port `6111`
- and lastly, the command to start the container is to invoke the binary created earlier, which is `./helloworld`

### Task #2 
Docker image:

- should have the name `go-helloworld`
- should have a valid tag, and a version with a major, minor, and patch included
- should be available in DockerHub, under your username e.g. `pixelpotato/go-helloworld`

### Task #3
Docker container:

- should be running on your local machine, by referencing the image from the DockerHub with a valid tag e.g. `pixelpotato/go-helloworld:v5.12.3`

**NOTE:** You will need to use `docker login` to login into Docker before pushing images to DockerHub.


## Solution

---

1. Create a Dockerfile manually or from the terminal: `touch Dockerfile`
2. Open the Dockerfile and create layers due to task #1 and save it:

```dockerfile
FROM golang:alpine

WORKDIR /go/src/app

ADD . .

RUN go build  -o helloworld

EXPOSE 6111

CMD ["./helloworld"]
```

3. Build the Docker image (make sure that you are in the go-helloworld directory):
`docker build -t go-helloworld .`

**NOTE:** If you get the permission denied response:
`Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post http://%2Fvar%2Frun%2Fdocker.sock/v1.24/build?buildargs=%7B%7D&cachefrom=%5B%5D&cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=Dockerfile&labels=%7B%7D&memory=0&memswap=0&networkmode=default&rm=1&shmsize=0&t=go-helloworld&target=&ulimits=null&version=1: dial unix /var/run/docker.sock: connect: permission denied`
   
Fix it:
- `sudo usermod -aG docker ${USER}`
- `su - ${USER}`
   
Check:
- `docker` ps -a
   
See details: [here](https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket)
   
Repeat: step 3 (build the Docker image)

**NOTE:** if you have error messages like this:
`go: go.mod file not found in current directory or any parent directory; see 'go help modules'` 
`The command '/bin/sh -c go build  -o helloworld' returned a non-zero code: 1`
it means you need to init `go.mod` file. We can handle this issue using one of the following approaches:
- add layer `RUN go mod init` before the `RUN go build  -o helloworld` to Dockerfile

1. Run the Docker image: `docker run -d -p 6111:6111 go-helloworld`
2. Verify if the application is available and work properly on localhost 127.0.0.1:6111 (if you have problems accessing `https`, just change it to `http`)
3. Tag the app: `docker tag go-helloworld <YOUR DOCKERHUB USER NAME>/go-helloworld:v1.0.0`
4. Push the app to the DockerHub repo (need account): `docker push <YOUR DOCKERHUB USER NAME>/go-helloworld:v1.0.0`
   
Feel free to check the [video Docker for Application Packaging](https://www.youtube.com/watch?v=f6gw_f-CO8U&t=1s)


## Exercise 2 Solution: Deploy Your First Kubernetes Cluster

#### run the application

- `kubectl run test --image pixelpotato/go-helloworld:v1.0.0`

#### access the application on the local host

- `kubectl port-forward test-97856cf4-6fvjw 7111:6111`


#### kubeconfig

- K3s stores the kubeconfig file under /etc/rancher/k3s/k3s.yaml path
- API server - https://127.0.0.1:6111
- authentication mechanism - username (admin) and password

#### kubectl commands

- `kubectl cluster-info` to get the control plane and add-ons endpoints
- `kubectl get nodes` - to get all the nodes in the cluster
- `kubectl get nodes -o wide` - to get extra details about the nodes, including internal IP
- `kubectl describe node node-name` - to get all the configuration details about the node, including the allocated pod CIDR

[Kubeconfig Walkthrough video tutorial](https://www.youtube.com/watch?v=yi1kR9nDw1g)
## Additional Resourses

[Organizing Cluster Access Using kubeconfig Files](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)
[Kubernetes Config file](https://community.suse.com/posts/cluster-this-is-your-admin-do-you-read)