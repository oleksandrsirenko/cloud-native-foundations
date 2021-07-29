# Exercise: Kubernetes Resources

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

Make sure the following tasks are compleated:

- [x] You have created a Namespace
- [x] You have created a Deployment
- [x] You have created a Service
- [x] You have created a Configmap
