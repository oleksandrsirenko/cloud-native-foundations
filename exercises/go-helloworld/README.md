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


## Excercise #1: Docker for Application Packaging

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

## Exercise #2: Deploy Your First Kubernetes Cluster

### Environment Setup

Provision a Vagrant box locally and install Kubernetes with k3s.

Note: Make sure to have VirtualBox `6.1.16` or higher installed.

- [x] Install [vagrant](https://www.vagrantup.com) on your machine
- [x] Clone the course repository using git commands
- [x] Navigate inside the `exercises` directory and examine the `Vagrantfile`
- [x] Create a Vagrant box using `vagrant up` command (Note: you need to be in the same directory as the Vagrantfile for this command to work)
- [x] SSH into the vagrant box by using `vagrant ssh` command
- [x] Deploy a Kubernetes cluster using [k3s](https://k3s.io/)
- [x] Examine your cluster using `kubectl` command(Note: you need to be root to access the kubeconfig file, use `sudo su -` before using kubectl commands)

### Exercise 

From the kubeconfig, identify:

- the IP and port of the API server
- authentication mechanism

Note: *Refer to the main [k3s installation guide](https://rancher.com/docs/k3s/latest/en/quick-start/#install-script), to find the location of the kubeconfig file.*

From the cluster using kubectl commands to identify:

- endpoints of the control plane and add-ons
- amount of nodes
- node internal IP
- the pod CIDR allocate to the node

Note: *To access kubectl commands you need to become root by using `sudo su -` command.*