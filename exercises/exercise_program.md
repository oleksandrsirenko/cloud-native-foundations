# Curriculum

1. Introduction to Cloud Native Fundamentals
2. Architecture Consideration for Cloud-Native Application
3. Container Orchestration with Kubernetes
4. Open Source PaaS
5. CI/CD with Cloud Native Tooling

## Exercise procedure

1. [Python Flask Application](python-helloworld/flask_app_exercise.md):
    1. Create endpoints for application status and metrics
    2. Implement application logging
    3. Dockerize python application (optional). Apply all the steps from exercise 2 to flask app: create Dokerfile, build and push Docker image to the Docker Hub)

2. [Docker for Application Packaging](go-helloworld/docker_exercise.md)
    1. Create Docker file for Go application
    2. Build Docker Image
    3. Push Doker image to Docker Hub

3. [Kubernetes](kubernetes/kubernetes_exercise.md)
    1. Create Kubernetes cluster
    2. Deploy application to the Kubernetes cluster
    3. Deploy Kubernetes resources using kubeclt command

4. [Dockerize Application with GitHub Actions](github-actions/github_actions_exercise.md)

5. [The CI/CD Fundamental](argocd/argocd_exercise.md)
    1. [Deploy the application using ArgoCD](argocd/argocd_exercise.md)
    2. [Create Helm Chart](helm/helm_exercise.md)
    3. [Deploy Helm Chart](helm/helm_exercise.md)

[<- Back to main README](../README.md)

[-> Explore solution list](../solutions/solution_list.md)