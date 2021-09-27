# Curriculum

1. Introduction to Cloud Native Fundamentals
2. Architecture Consideration for Cloud-Native Application
3. Container Orchestration with Kubernetes
4. Open Source PaaS
5. CI/CD with Cloud Native Tooling

## Solutions

1. Python Flask Application: `solutions/python-helloworld/`
    1. Create endpoints for application status and metrics: `endpoints.md`
    2. Implement application logging: `logging.md`
    3. Dockerize python application (optional). Apply all the steps from exercise 2 to flask app: create Dokerfile, build and push Docker image to the Docker Hub): `dokerize_app.md`

2. Docker for Application Packaging: `solutions/go-helloworld/`
    1. Create Docker file for Go application: `docker_packaging.md`
    2. Build Docker Image: `docker_packaging.md`
    3. Push Doker image to Docker Hub: `docker_packaging.md`

3. Kubernetes: `solutions/kubernetes/`
    1. Create Kubernetes cluster: `k3s_cluster.md`
    2. Deploy application to the Kubernetes cluster: `k3s_cluster.md`
    3. Deploy Kubernetes resources using kubeclt command: `kubernetes_resources.md`

4. Dockerize Application with GitHub Actions: `solutions/github-actions/build_push_actions.md`

5. The CI/CD Fundamental
    1. Deploy the application using ArgoCD: `solutions/argocd/cd_with_argocd.md`
    2. Create Helm Chart: `solutions/helm/config_managers.md`
    3. Deploy Helm Chart: `solutions/helm/config_managers.md`
