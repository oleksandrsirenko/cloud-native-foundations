# :bulb: Solution 1.3: Build a Docker Image and Push It to the Docker Hub

- install docker if not have it yet: `sudo apt-get install docker.io`
- create a Dockerfile like this (change maintainer name):

```dockerfile
FROM python:3.8
LABEL maintainer="Oleksandr Sirenko"

COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt

# command to run on container start
CMD [ "python", "app.py" ]
```

- `docker build -t python-helloworld .` - build a docker image
- `docker tag go-helloworld <YOUR DOCKERHUB USER NAME>/go-helloworld:v1.0.0` - tag your app if you need
- `docker push <YOUR DOCKERHUB USER NAME>/go-helloworld:v1.0.0` - push the app to the DockerHub repo (need account)

[-> Back to solution list](../solution_list.md)
