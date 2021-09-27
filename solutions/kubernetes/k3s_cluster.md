# :pencil2: Exercise 3: Deploy Your First Kubernetes Cluster

In the previous steps, we have already created simple applications, dockerized them, and pushed Doker images to the docker hub. So the next step will be to deploy one of these apps from image to Kubernetes cluster on our local machine directly from our Docker Hub.

> :pushpin: **NOTE: This exercise is slightly different from the original one - it was modified according to more practical use cases and connect student progress from the previous lessons.**

## Exercise 3.1: Create Kubernetes Cluster

Create k3s Kubernetes cluster on the local machine using Virtual Box and Vagrant

## Exercise 3.2: Deploy Application to the Kubernetes Cluster

Deploy an application of your choice (go-helloworld or python-helloworld) to the cluster directly from the Docker Hub.

## :bulb: Solution 3.1: Create Kubernetes Cluster

### Install Required Tools

- `sudo apt install virtualbox` - install VirtualBox this is the prerequisites to using Vagrant
- `sudo apt-get update && sudo apt-get install vagrant` - install Vagrant

### Up Virtual Machine

- `cd exercises` - change directory to exercises (where the Vagrant file is)
- `vagrant up` - up the virtual box, if necessary init with the `vagrant init` first
- `vagrant status` - check the status of Virtual Machine

> :bulb: **Useful commands**: 

- run `vagrant halt` to shut down Virtual Machine forcefully;
- run `vagrant suspend` to simply suspend the Virtual Machine;
- run `vagrant up` to restart Virtual Machine again.

## Start Kubernetes Cluster

- `vagrant ssh` - next go to the vagrant shell
- `curl -sfL https://get.k3s.io | sh` get k3s with one command in vagrant environment (inside the shell) see details on [k3s.io](https://k3s.io/)
- check nodes `kubectl get no` if get an error like permission denied need to configure the `/etc/rancher/k3s/k3s.yaml` or use `sudo su` - a command to perform as a superuser.

> :bulb: **Useful commands**: 

- stop the Kubernetes cluster, as the root user: `shutdown -h now`

## Solution 3.2: Deploy Application to the Kubernetes Cluster

To run the app at a cluster use the command:

- `kubectl run <POD NAME> --image=<DOCKER IMAGE PATH>` e.g., `kubectl run test --image=pixelpotato/go-helloworld:v1.0.0`

> :pushpin: **NOTE: You need to check pod status with** `kubectl get pods` **or** `kubectl describe pod <POD-NAME>`. **If the status of the reated pod is** `CreateContainerError` **you need to fix this issue with the command:** `zypper install -t pattern apparmor`

Use the following kubectl command syntax to deploy the app to the cluster directly from your Docker Hub:

- `kubectl create deployment <DEPLOYMENT NAME> --image=docker.io/<DOCKERHUB USERNAME>/<DOKER IMAGE NAME>:<TAG>` e.g., `kubectl create deployment test-go  --image=docker.io/pixelpotato/go-helloworld:v1.0.0`

### Access the Application on the Local Host

- `kubectl port-forward <POD NAME> 6111:6111`

### Kubeconfig

- K3s stores the kubeconfig file under /etc/rancher/k3s/k3s.yaml path
- API server - <https://127.0.0.1:6111>
- authentication mechanism - username (admin) and password

### Useful kubectl commands (more in [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/))

- `kubectl cluster-info` to get the control plane and add-ons endpoints
- `kubectl get nodes` - to get all the nodes in the cluster
- `kubectl get nodes -o wide` - to get extra details about the nodes, including internal IP
- `kubectl describe node <NODE NAME>` - to get all the configuration details about the node, including the allocated pod CIDR
- `kubectl get pods` - basic pods information
- `kubectl describe pod <POD NAME>` - to check the detailed information of a particular pod
- `kubectl delete pod <POD NAME>` - to delete pod

[Kubeconfig Walkthrough video tutorial](https://www.youtube.com/watch?v=yi1kR9nDw1g)

## Additional Resourses

1. [Stopping and starting Kubernetes cluster](https://www.ibm.com/docs/en/fci/1.0.3?topic=kubernetes-stopping-starting-cluster)
2. [Organizing Cluster Access Using kubeconfig Files](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)
3. [How to Manage Kubernetes With Kubectl](https://rancher.com/learning-paths/how-to-manage-kubernetes-with-kubectl/)
4. [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
5. [Kubernetes Config file](https://community.suse.com/posts/cluster-this-is-your-admin-do-you-read)
6. [Using kubectl to Create a Deployment](https://kubernetes.io/docs/tutorials/kubernetes-basics/deploy-app/deploy-intro/)
7. [How to Delete Pods from a Kubernetes Node](https://www.bluematador.com/blog/safely-removing-pods-from-a-kubernetes-node)
8. [Use Port Forwarding to Access Applications in a Cluster](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/)

[-> Back to solution list](../solution_list.md)