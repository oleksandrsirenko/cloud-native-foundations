# Configuration Managers

For the management of multiple declarative Kubernetes manifests, a templating layer is necessary, especially if the application is replicated across different regions. For this purpose configuration managers, such as Helm and Kustomize, were introduced.

This exercise will focus on creating your first Helm chart to deploy multiple Nginx applications using the same template and multiple input files.

## Exercises 5.2-5.3: Helm with ArgoCD

Using the manifests provided in the course repository, create a helm chart (Chart.yaml, templates, values.yaml) that will template the following parameters:

- namespace name
- replica count
- image:
  - name
  - tag
  - pull policy
- resources
  - requests for CPU and memory
- service
  - port
  - type (e.g. ClusterIP)
- configmap data (e.g. the key-value pair)

The chart details should be as following:

- name: nginx-deployment
- version: 1.0.0
- keywords: nginx

Once the Helm chart is available make sure that a default `values.yaml` file is available. This values file will be used as a default input file for the Helm chart. The `values.yaml` file should have the following specification:

- **values.yaml**
  - namespace name: demo
  - replica count: 3
  - image repository: nginx
  - image tag: alpine
  - image pull policy: IfNotPresent
  - resources: CPU 50m and memory 256Mi
  - service type: ClusterIP
  - service port: 8111
  - configmap data: "version: alpine"

Next, create 2 values files with the following specifications:

- **values-staging.yaml**
  - namespace name: staging
  - replica count: 1
  - image repository: nginx
  - image tag: 1.18.0
  - resources: CPU 50m and memory 128Mi
  - configmap data: "version: 1.18.0"

- **values-prod.yaml**
  - namespace name: prod
  - replica count: 2
  - image repository: nginx
  - image tag: 1.17.0
  - resources: CPU 70m and memory 256Mi
  - service port: 80
  - configmap data: "version: 1.17.0"

Finally, using the values files above (values-prod, values-staging), create 2 ArgoCD application, `nginx-staging` and `nginx-prod` respectively. These should deploy the nginx Helm Chart referencing each input values files.

## Solutions 5.2-5.3: Helm with ArgoCD

Helm provides a powerful mechanism to inject values to templated YAML manifests.

### 5.2: Create Helm Chart

The full Helm chart for `nginx-deployment` can found in the course repository.

The Helm chart is defined in the `Chart.yaml` file, which contains the API version, name and version of the chart:

```yaml
apiVersion: v1
name: nginx-deployment
description: Install Nginx deployment manifests 
keywords:
- nginx 
version: 1.0.0
maintainers:
- name: kgamanji 
```

An example of the values.yaml file can be found below:

```yaml
namespace:
  name: demo

service:
  port: 8111
  type: ClusterIP

image:
  repository: nginx 
  tag: alpine
  pullPolicy: IfNotPresent

replicaCount: 3

resources:
  requests:
    cpu: 50m
    memory: 256Mi

configmap:
  data: "version: alpine"
```

The above configuration represents the default parameters of application deployment if it is not overwritten by a different values file.

Below is an example of the `values-prod.yaml` file, which will override the default parameters:

```yaml
namespace:
  name: prod 

service:
  port: 80
  type: ClusterIP

image:
  repository: nginx 
  tag: 1.17.0
  pullPolicy: IfNotPresent

replicaCount: 2

resources:
  requests:
    cpu: 70m
    memory: 256Mi

configmap:
  data: "version: 1.17.0"
```

The `values-staging.yaml` can be found in the course repository `solutions/helm/nginx-deployment`

### 5.3: Deploy Helm Chart

And finally, here is ArgoCD application CRD for the `nginx-prod` deployment:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: nginx-prod
  namespace: argocd
spec:
  destination:
    namespace: default
    server: https://kubernetes.default.svc
  project: default
  source:
    helm:
      valueFiles:
      - values-prod.yaml
    path: helm
    repoURL: https://github.com/kgamanji/argocd-demo
    targetRevision: HEAD
```

The `nginx-staging.yaml` Application for ArgoCD can be found in the course repository.

Additional resources:

1. [Helm Walkthrough video](https://www.youtube.com/watch?v=i1ZvdS7qdAI)
2. [Create Helm Chart](https://www.youtube.com/watch?v=aDqyHi2iiQk)
3. [Deploy Helm Chart](https://www.youtube.com/watch?v=k-7N7gcwW5Y)

[-> Back to solution list](../solution_list.md)
