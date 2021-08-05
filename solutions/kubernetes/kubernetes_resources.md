# Exercise 3.3: Kubernetes Resources

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

**Note**: Nginx is one of the public Docker images, that you can access and use for your exercises or testing purposes.

Make sure the following tasks are completed:

- [x] You have created a Namespace
- [x] You have created a Deployment
- [x] You have created a Service
- [x] You have created a Configmap

## Solution 3.3: Kubernetes Resources

### Create the namespace

**Note**: label option is not available with `kubectl create`

- `kubectl create ns demo`

### Label the namespace

- `kubectl label ns demo tier=test`

### Create the nginx-alpine deployment

- `kubectl create deploy nginx-alpine --image=nginx:alpine  --replicas=3 --namespace demo`

### Label the deployment

- `kubectl label deploy nginx-alpine app=nginx tag=alpine --namespace demo`

### Expose the nginx-alpine deployment, which will create a service

- `kubectl expose deployment nginx-alpine --port=8111 --namespace demo`

### Create a config map

- `kubectl create configmap nginx-version --from-literal=version=alpine --namespace demo`
