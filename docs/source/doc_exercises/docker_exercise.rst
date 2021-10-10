2. Docker for Application Packaging
===================================

.. |2.0_docker_overview| raw:: html

  <div class="video-container">
    <iframe src="https://www.youtube.com/embed/iqqDU2crIEQ" title="GitHub Actions - Supercharge your GitHub Flow" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
  </div>

|2.0_docker_overview|

In this exercise, you will use a minimal Go application that looks like this:

.. code-block:: go
   
   package main

   import (
	   "fmt"
	   "net/http"
   )

   func helloWorld(w http.ResponseWriter, r *http.Request) {
	   fmt.Fprintf(w, "Sherlock Holmes sat silently...")
   }

   func main() {
	   http.HandleFunc("/", helloWorld)
	   http.ListenAndServe(":8080", nil)
   }


Your tasks are to containerize this Go application and push it to Docker Hub. But 
before you get started, we highly recommend watching the `Docker introductory video <https://www.youtube.com/embed/iqqDU2crIEQ>`__ at the 
top of the page and following through with the `How to Build a Containerized Go 
Application with Docker <https://docs.docker.com/language/golang/>`__ tutorial. 

.. _lesson_2_prep:

Preparation
-----------

.. note::
   This application listens on port ``8080``

1. `Download and install Docker Engine <https://docs.docker.com/engine/install/>`__
2. `Install Go <https://golang.org/doc/install>`__: ``sudo apt install golang-go``
3. Change the working directory: ``cd exercises/go-helloworld``
4. Run Go app using the terminal command: ``go run main.go``
5. Check your app in a web browser at: http://127.0.0.1:8080/

Exercise 2.1 - Create Dockerfile
--------------------------------

Dockerfile requirements:

-  Use an image that is based on the latest stable Go environment.
-  Set the proper working directory.
-  Copy source code into the image.
-  Implement logic to build the application.
-  Make sure that app access on a default port ``8080``.
-  Add the command to start the container.

Exercise 2.2 - Build a Docker Image
-----------------------------------

Docker image requirements:

-  Use the ``go-helloworld`` name as the image name.
-  Tag the image to ``YOUR_DOCKERHUB_USERNAME/go-helloworld:v1.0.0``, 
   (don't forget to replace ``YOUR_DOCKERHUB_USERNAME`` 
   with your Docker Hub username).


Exercise 2.3 - Push a Docker Container to Docker Hub
----------------------------------------------------

.. note:: 
   To push a Docker container image to Docker Hub,
   you need to have `a Docker Hub account <https://hub.docker.com/>`__

The application image shoud be available on Docker Hub under the 
``https://hub.docker.com/r/YOUR_DOCKERHUB_USERNAME/go-helloworld``, 
where ``YOUR_DOCKERHUB_USERNAME`` is you current user name 
e.g. https://hub.docker.com/r/osdoc/go-helloworld.


2.4 Pull Image From the Docker Hub
----------------------------------

To complete the exercise, do the tasks in the following order:

1. Remove all containers and images from your local computer.
2. Pull the image from the Docker Hub repository.
3. Run the pulled image on your local computer.
4. Verify go-helloworld application at: http://127.0.0.1:8080/

Exercise 2.5 (Optional) - Containerize Python Application
---------------------------------------------------------

This is an exercise to consolidate the material covered in the section. 
Your job is to containerize the Flask application that you created in 
the first lesson.

Follow all these steps to reach your goal:

1. Create Dockerfile.
2. Build a Docker image.
3. Run newly built image as a container.
4. Check the Flask application at http://127.0.0.1:5000/
5. Push the Docker container to Docker Hub.
6. Remove container from the local machine.
7. Pull the container from Docker Hub.
8. Run pulled container.


Additional Resources
--------------------

1. `Docker Tutorial for Beginners - 3 Hour Video Course <https://youtu.be/3c-iBn73dDE>`__
2. `Docker Hub Quickstart <https://docs.docker.com/docker-hub/>`__
3. `Build a Docker Image <https://docs.docker.com/engine/reference/commandline/build/>`__
4. `Pushing a Docker Container Image to Doker Hub <https://docs.docker.com/docker-hub/repos/#pushing-a-docker-container-image-to-docker-hub>`__
5. `How to Build a Containerized Go Application with Docker <https://docs.docker.com/language/golang/>`__
6. `Docker Image Pipeline for Go <https://codefresh.io/docs/docs/learn-by-example/golang/golang-hello-world/>`__
7. `Build Python Docker Image <https://docs.docker.com/language/python/build-images/>`__
8. `How to Serve a Flask App with Amazon Lightsail Containers <https://aws.amazon.com/getting-started/hands-on/serve-a-flask-app/>`__
9. `Docker Cheet Sheet <https://www.docker.com/sites/default/files/d8/2019-09/docker-cheat-sheet.pdf>`__
10. `How To Install and Use Docker on Ubuntu 20.04 <https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04>`__
11. `A Beginner-Friendly Introduction to Containers, VMs and Docker <https://www.freecodecamp.org/news/a-beginner-friendly-introduction-to-containers-vms-and-docker-79a9e3e119b/>`__