2. Docker for Application Packaging
===================================

.. note:: 
   Make sure you have :ref:`prepared the environment <lesson_2_prep>` for this task.


Solution 2.1 - Create Dockerfile
--------------------------------

1. Create Dockerfile: ``touch Dockerfile``
2. Open the Dockerfile.
3. Create layers due to task:

.. code-block:: dockerfile

   # syntax=docker/dockerfile:1

   FROM golang:1.16-alpine

   WORKDIR /go/src/app

   ADD . .

   RUN go mod init

   RUN go build -o go-helloworld

   EXPOSE 8080

   CMD ["./go-helloworld"]

4. Save changes

Solution 2.2 - Build a Docker Image
-----------------------------------

1. Make sure you are in the ``/exercises/go-helloworld/`` directory.
2. Build a Docker image using the prompt command:

.. code-block:: shell

   docker build -t go-helloworld .


Solution 2.3 - Push a Docker Container to Docker Hub
----------------------------------------------------

.. note:: 
   To push a Docker container to Docker Hub,
   you need to have `a Docker Hub account <https://docs.docker.com/docker-hub/>`__

Follow these steps to push a Docker container to Docker Hub:

1. Run a Docker image as a container: ``docker run -p 8080:8080 --name go_moriarty -d go-helloworld``
2. Verify if the application at: http://127.0.0.1:8080/
3. Login to the Docker Hub: ``docker login -u "YOUR_DOCKERHUB_USERNAME" -p "YOUR_DOCKERHUB_PASSWORD" docker.io``
4. Tag the image: ``docker tag go-helloworld YOUR_DOCKERHUB_USERNAME/go-helloworld:v1.0.0``
5. Push the image to the DockerHub repo: ``docker push YOUR_DOCKERHUB_USERNAME/go-helloworld:v1.0.0``

Feel free to check the `video Docker for Application Packaging <https://www.youtube.com/watch?v=f6gw_f-CO8U&t=1s>`__

.. |2.3_docker_hub_video| raw:: html

    <div class="video-container">
      <iframe src="https://www.youtube.com/embed/f6gw_f-CO8U" title="Push a Docker Image to Docker Hub" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    </div>

|2.3_docker_hub_video|


2.4 Pull Image From the Docker Hub
----------------------------------

Remove all builds and images:

- Explore the processes: ``docker ps -a``
- Stop the container: ``docker stop go_moriarty``
- Remove the container: ``docker rm go_moriarty``
- Explore the images: ``docker images``
- Remove the image from local machine: ``docker rmi YOUR_DOCKERHUB_USERNAME/go-helloworld:v1.0.0``

Pull container from your Docker Hub repository:

- Pull the image: ``docker pull YOUR_DOCKERHUB_USERNAME/go-helloworld:v1.0.0``
- Run the image as a container: ``docker run -p 8080:8080 --name go_pulled -d YOUR_DOCKERHUB_USERNAME/go-helloworld:v1.0.0``
- Verify go-helloworld application at: http://127.0.0.1:8080/


Useful Docker Commands
~~~~~~~~~~~~~~~~~~~~~~

- Get container name: ``docker ps -a``
- Stop the container: ``docker stop CONTAINER_NAME``
- Remove the container: ``docker rm CONTAINER_NAME``
- Stop all the processes: ``docker kill $(docker ps -q)``
- Remove all the processes: ``docker rm $(docker ps -a -q)``
- Remove all the images: ``docker rmi $(docker images -q)``

Find out more helpful commands in the `Docker Cheat Sheet. <https://www.docker.com/sites/default/files/d8/2019-09/docker-cheat-sheet.pdf>`__ 


Solution 2.5 - Dockerize Python Flask Application
-------------------------------------------------

In this step, we have gained a basic understanding of containerizing applications with Docker. 
So let's take a look at the process of dockerizing a Pyhton Flask application.

Use this recipe to put your Python Flask application in a Docker container and push it to the Docker Hub:

1. Install Docker if not have it yet: ``sudo apt-get install docker.io``
2. Activate Python virtual environment: ``source venv/bin/activate``
3. Change working directory: ``cd exercise/python-helloworld``
4. Create a Dockerfile:

.. code-block:: dockerfile

   # syntax=docker/dockerfile:1

   FROM python:3.8

   WORKDIR /app

   COPY . /app

   RUN pip install -r requirements.txt

   EXPOSE 5000

   CMD [ "python", "app.py" ]

3. Run the Docker image: ``docker run -p 5000:5000 --name monty_python -d python-helloworld``
4. Verify if the application is available and work properly at: http://127.0.0.1:5000/
5. Build a Docker image: ``docker build -t python-helloworld .`` 
6. Tag application: ``docker tag python-helloworld YOUR_DOCKERHUB_USERNAME/python-helloworld:v1.0.0``
7. Login to Docker Hub: ``docker login``
8. Push the application to Docker Hub: ``docker push YOUR_DOCKERHUB_USERNAME/python-helloworld:v1.0.0``

Common Errors & How to Fix Them
-------------------------------

Permission Denied Error
~~~~~~~~~~~~~~~~~~~~~~~

.. error:: ``Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post http://%2Fvar%2Frun%2Fdocker.sock/v1.24/build?buildargs=%7B%7D&cachefrom=%5B%5D&cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=Dockerfile&labels=%7B%7D&memory=0&memswap=0&networkmode=default&rm=1&shmsize=0&t=go-helloworld&target=&ulimits=null&version=1: dial unix /var/run/docker.sock: connect: permission denied```

To fix **Permission Denied Error**, follow the next steps:

1. Create the docker group: ``sudo groupadd docker``
2. Add your user to the docker group: ``sudo usermod -aG docker ${USER}``
3. Set the superuser: ``su -s ${USER}``
4. Check if Docker work correctly: ``docker ps -a``
5. Build Docker image: ``docker build -t go-helloworld .``

Follow the `link <https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket>`__ to learn more about how to fix **Permission Denied Error**.

File Not Found Error
~~~~~~~~~~~~~~~~~~~~

.. error:: 
   ``go: go.mod file not found in current directory or any parent directory; see 'go help modules'``

This error refers to building a Docker image for the Golang application and means 
you need to init ``go.mod`` file. To fix **File Not Found Error**, follow these steps:

1. Open the Dockerfile of your Go application.
2. Add ``RUN go mod init`` layer before the ``RUN go build -o helloworld`` layer.
3. Save the changes.
4. Build Docker image using ``docker build -t go-helloworld .`` command.

After these manipulations, the Dockerfile should look like this:

.. code-block:: dockerfile
   :emphasize-lines: 7

   FROM golang:alpine

   WORKDIR /go/src/app

   ADD . .

   RUN go mod init

   RUN go build -o helloworld

   EXPOSE 8080

   CMD ["./helloworld"]

