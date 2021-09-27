# Exercise 5: The CD Fundamentals

Continuous Delivery (CD) is the ability to get code changes reliably to production environments. This practice should be automated and should enable developers to provide value to consumers efficiently.

In this exercise, you will use ArgoCD to automate the delivery of an application to a Kubernetes cluster.

## Environment Setup

Prepare the environment to create your delivery pipeline.

Note: Make sure to have [VirtualBox](https://www.virtualbox.org/wiki/Downloads) installed.

- [x] Install vagrant on your machine
- [x] Examine the Vagrantfile in the course repository
- [x] Create a vagrant box by using `vagrant up` command (Note: you need to be in the same repository as the Vagrant file for this command to work)
- [x] SSH into the vagrant box by using `vagrant ssh` command
- [x] Deploy a Kubernetes cluster using k3s
- [x] Examine you cluster using `kubectl` command (Note: you need to be root to access the kubeconfig file, use `sudo su -` before using kubectl commands)

## Exercise 5.1: Deploy the Application Using ArgoCD

- [x] Install ArgoCD by referring to the [installation guide](https://argoproj.github.io/argo-cd/getting_started/)
- [x] Expose the `argocd-server` using a `NodePort` service on port `30007 `on HTTP and `30008` on HTTPS. The YAML manifest for the NodePost service can be found in the course repository
- [x] Access the ArgoCD UI by going to `https://192.168.50.4.30008` or `https://192.168.50.4.30007`
- [x] Insert login credantials by using credentials guide
- [x] Create an ArgoCD application named `nginx-alpine` and use the manifests provided in the course repository

[-> Back to exercise list](../exercise_program.md)
