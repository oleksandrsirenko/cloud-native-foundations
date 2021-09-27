# Exercise 3: Deploy Your First Kubernetes Cluster

In the previous steps, we have already created simple applications, dockerized them, and pushed Doker images to the docker hub. So the next step will be to deploy one of these apps from image to Kubernetes cluster on our local machine directly from our Docker Hub.

> :pushpin: **NOTE: This exercise is slightly different from the original one - it was modified according to more practical use cases and connect student progress from the previous lessons.**

## Task 1: Create Kubernetes Cluster

Create k3s Kubernetes cluster on the local machine using Virtual Box and Vagrant

## Task 2: Deploy Application to the Kubernetes Cluster

Deploy an application of your choice (go-helloworld or python-helloworld) to the cluster directly from the Docker Hub.

## Task 3: Kubernetes Resources

Now you have learned many Kubernetes recourses, in this exercise, you will deploy the following resources using the kubectl command.

- a namespace
  - name: `demo`
  - label: `tier: test`
- a deployment:
  - image: `nginx:alpine`
  - name:`nginx-apline`
  - namespace: `demo`
  - replicas: 3
  - labels: `app: nginx`, `tag: alpine`
- a service:
  - expose the above deployment on port `8111`
  - namespace: `demo`
- a configmap:
  - name: `nginx-version`
  - containing key-value pair: `version=alpine`
  - namespace: `demo`

> :pushpin: **Note: Nginx is one of the public Docker images, that you can access and use for your exercises or testing purposes.**

Make sure the following tasks are completed:

- [x] You have created a Namespace
- [x] You have created a Deployment
- [x] You have created a Service
- [x] You have created a Configmap

[-> Back to exercise list](../exercise_program.md)