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
