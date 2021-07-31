# Exercise: The CD Fundamentals

Continuous Delivery (CD) is the ability to get code changes reliably to production environments. This practice should be automated and should enable developers to provide value to consumers efficiently.

In this exercise, you will use ArgoCD to automate the delivery of an application to a Kubernetes cluster.

## Environment Setup

Prepare the environment to create your delivery pipeline.

Note: Make sure to have [VirtualBox](https://www.virtualbox.org/wiki/Downloads) installed.

Install VirtualBox and Vagrant:

- `sudo apt install virtualbox` - install VirtualBox this is the prerequisites to using Vagrant
- `sudo apt-get update && sudo apt-get install vagrant` - install Vagrant

Up Virtual Machine:

- `cd exercises` - change directory to exercises (where the Vagrant file is)
- `vagrant up` - up the virtual box, if nessesary init with the `vagrant init` first
- `vagrant status` - check the status of virtual machine

**Notes**: *To stop this VM, you can run `vagrant halt` to shut it down forcefully, or you can run `vagrant suspend` to simply suspend the virtual machine. In either case, to restart it again, simply run `vagrant up`.

Kubernetes Cluster:

- `vagrant ssh` - next go to the vagrant shell
- `curl -sfL https://get.k3s.io | sh` get k3s with one command in vagrant environment (inside the shell) see details on [k3s.io](https://k3s.io/)
- check nodes `kubectl get no` if get error like permission denied need to configure the `/etc/rancher/k3s/k3s.yaml` or use `sudo su` - a command to perform as a superuser.

**Notes**:

- `kubectl config view` - to view a configuration of cluster
- `shutdown -h now` - to stop the Kubernetes cluster, as the root user

## Solution: Deploy application using ArgoCD

To [install ArgoCD](https://argoproj.github.io/argo-cd/getting_started/#1-install-argo-cd) use following commands inside the vagrant shell:

- `kubectl create namespace argocd`
- `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`

Then check pods according to the created argocd namespace (You can see five runnig containers and their status):

- `kubectl get po -n argocd`

Next step we need to expose our application to be accessible from the vagrant box to our host machine.

- `kubectl get svc -n argocd` get all the services
- `vim argocd-server-nodeport.yaml` - create a file within the shell and write the manifest:

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
  labels:
    app.kubernetes.io/component: server
    app.kubernetes.io/name: argocd-server
    app.kubernetes.io/part-of: argocd
  name: argocd-server-nodeport
  namespace: argocd
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 8080
    nodePort: 30007
  - name: https
    port: 443
    protocol: TCP
    targetPort: 8080
    nodePort: 30008
  selector:
    app.kubernetes.io/name: argocd-server
  sessionAffinity: None
  type: NodePort
```

`esc` and `:wq` to save changes

Next step is apply created yaml manifest with the following command:

- `kubectl apply -f argocd-server-nodeport.yaml` 

Then in the same manner create the manifest for application

- `vim nginx-alpine.yaml`

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: nginx-alpine
  namespace: argocd
spec:
  destination:
    namespace: default
    server: https://kubernetes.default.svc
  project: default
  source:
    path: exercises/manifests 
    repoURL: https://github.com/udacity/nd064_course_1 
    targetRevision: HEAD
  # Sync policy
  syncPolicy: {}
  ```

`esc` and `:wq` to save changes

- `kubectl apply -f nginx-alpine.yaml` - appling will create application with the command line messege: `application.argoproj.io/nginx-alpine created`

- `kubectl get application -n argocd` - check the app

**NOTE:** Pay attention when copy/past manifests - often letter `a` in the first string `apiVersion` lost and it look like `piVersion` that cause the error like this: 

`error: error validating "argocd-server-nodeport.yaml": error validating data: apiVersion not set; if you choose to ignore these errors, turn validation off with --validate=false`

So be careful and check if everything was written correctly, e.g.:

- `cat nginx-alpine.yaml`

## Use ArgoCD in your browser at

- http://192.168.50.4.30007
- http://192.168.50.4.30008

## Additional Resources

1. [Guide To GitOps](https://www.weave.works/technologies/gitops/)
2. [ArgoCD Installation video](https://www.youtube.com/watch?v=TJrSM31Jj_8)
3. [Deploy app with ArgoCD](https://www.youtube.com/watch?v=mYg-ULq9Rzg)