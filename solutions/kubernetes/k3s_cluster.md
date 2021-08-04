# Exercise: Deploy Your First Kubernetes Cluster

---

In the previous steps we have already created simple applications, dokerized them and pushed Doker images to the docker hub. So the next step will be to deploy one of this apps from image to kubernetes cluster on our local machine directly from our docker hub. 

**NOTE**: This exercise is slightly different from the original one - it was modifited according to more practical use case and connect student progress from the previous lessons.

Task 1: Create k3s kubernates cluster on the local machine using Virtual Box and Vagrant 
Task 2: Deploy application of your choise (go-helloworld or python-helloworld) to the cluster directly from the .

## Task 1 Solution: Create Kubernetes Cluster

### Install Required Tools

- `sudo apt install virtualbox` - install VirtualBox this is the prerequisites to using Vagrant
- `sudo apt-get update && sudo apt-get install vagrant` - install Vagrant

### Up Virtual Machine

- `cd exercises` - change directory to exercises (where the Vagrant file is)
- `vagrant up` - up the virtual box, if nessesary init with the `vagrant init` first
- `vagrant status` - check the status of virtual machine

**NOTE**: *To stop this VM, you can run `vagrant halt` to shut it down forcefully, or you can run `vagrant suspend` to simply suspend the virtual machine. In either case, to restart it again, simply run `vagrant up`.

## Start Kubernetes Cluster

- `vagrant ssh` - next go to the vagrant shell
- `curl -sfL https://get.k3s.io | sh` get k3s with one command in vagrant environment (inside the shell) see details on [k3s.io](https://k3s.io/)
- check nodes `kubectl get no` if get error like permission denied need to configure the `/etc/rancher/k3s/k3s.yaml` or use `sudo su` - a command to perform as a superuser.

**NOTE**: To stop the Kubernetes cluster, as the root user, enter the following command: `shutdown -h now`

## Deploy Application to the Kubernetes Cluster

To run the app at a cluster use the command:

- `kubectl run <POD NAME> --image=<DOCKER IMAGE PATH>` e.g., `kubectl run test --image=pixelpotato/go-helloworld:v1.0.0`

**NOTE**: You need to check pod status with `kubectl get pods` or `kubectl describe pod <POD-NAME>`. If the status of created pod is `CreateContainerError` you need to fix this issue with command: `zypper install -t pattern apparmor`

To deploy the app to the cluster directly from your Docker hub kubectl command syntax:

- `kubectl create deployment <DEPLOYMENT NAME> --image=docker.io/<DOCKERHUB USERNAME>/<DOKER IMAGE NAME>:<TAG>` e.g. `kubectl create deployment test-go  --image=docker.io/pixelpotato/go-helloworld:v1.0.0`

### Access the Application on the Local Host

- `kubectl port-forward <POD NAME> 7111:6111`

- [] **NEED TO TEST IT ACCESS AGAIN** An issue with access to the app at the localhost

### Kubeconfig

- K3s stores the kubeconfig file under /etc/rancher/k3s/k3s.yaml path
- API server - https://127.0.0.1:6111
- authentication mechanism - username (admin) and password

### Useful kubectl commands (more in [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/))

- `kubectl cluster-info` to get the control plane and add-ons endpoints
- `kubectl get nodes` - to get all the nodes in the cluster
- `kubectl get nodes -o wide` - to get extra details about the nodes, including internal IP
- `kubectl describe node <NODE NAME>` - to get all the configuration details about the node, including the allocated pod CIDR
- `kubectl get pods` - basic pods information
- `kubectl describe pod <POD NAME>` - to check the detaled information of a particular pod
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