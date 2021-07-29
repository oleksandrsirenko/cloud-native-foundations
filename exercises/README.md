# Learing Schema

---

1. Introdiction to Cloud Native Fundamentals
2. Architecture Consideration for Cloud Native Application 
3. Container Orchestraition with Kubernetes
4. Open Source PaaS
5. CI/CD with Cloud Native Tooling

TODO GLOBAL:
- Refine the structure: add the tasks for each lesson in the exercises folder, transfer all solutions READMEs with tips to solutions folder
- Maybe even add the Learning Schedule file in the root of repository
- Add `How to Use` to main README file with the instruction like: look into the task description and try to solve, if you have a problem go to solutions and compare or find a tips how to handle errors

# Lesson 2 Architecture Consideration for Cloud Native Application

TODO: 
- Transfer the solution from README in exercise to solutions. The exercises must be clean, at least this is a good idea for the education purposes

# Lesson 3 - Container Orchestraition with Kubernetes

TODO 1:
- Create all tasks for Kubernetes Deployment [see demo](https://www.youtube.com/watch?v=8Dh3Z2lBR3s)
- Step by step solution guide, including installation of the required packages and tools
- Add external resources to syntax, guides, docs
- Test the solution one more time
- Look at the discussion and pull the most common issues
- Update task, add it to main pool and commit solution

TODO 1 PROGRESS:

Install

- `sudo apt install virtualbox` - install VirtualBox this is the prerequisites to using Vagrant
- `sudo apt-get update && sudo apt-get install vagrant` - install Vagrant

Up Virtual Machine

- `cd exercises` - change directory to exercises (where the Vagrant file is)
- `vagrant up` - up the virtual box, if nessesary init with the `vagrant init` first
- `vagrant status` - check the status of virtual machine

**Notes**: *To stop this VM, you can run `vagrant halt` to shut it down forcefully, or you can run `vagrant suspend` to simply suspend the virtual machine. In either case, to restart it again, simply run `vagrant up`.

Kubernetes Cluster

- `vagrant ssh` - next go to the vagrant shell
- `curl -sfL https://get.k3s.io | sh` get k3s with one command in vagrant environment (inside the shell) see details on [k3s.io](https://k3s.io/)
- check nodes `kubectl get no` if get error like permission denied need to configure the `/etc/rancher/k3s/k3s.yaml` or use `sudo su` - a command to perform as a superuser.

**Notes**:
1. To view a configuration of cluster:
`kubectl config view`
2. To stop the Kubernetes cluster, as the root user, enter the following command:
`shutdown -h now`

sources to check:
1. [Stopping and starting Kubernetes cluster](https://www.ibm.com/docs/en/fci/1.0.3?topic=kubernetes-stopping-starting-cluster)
2. [Organizing Cluster Access Using kubeconfig Files](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)
3. [How to Manage Kubernetes With Kubectl](https://rancher.com/learning-paths/how-to-manage-kubernetes-with-kubectl/)

## Lesson 3 part 2 (16-27)
  - Kubernetes Resources
  - Kubernetes Manifests

item 16. typos "A `note` can host multiple pods for different applications." `noDe`