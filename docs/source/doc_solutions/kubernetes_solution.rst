3. Kubernetes Cluster
=====================

.. note:: 
   Make sure you have :ref:`prepared the environment <lesson_3_prep>` for this task.


Solution 3.1 - Create Kubernetes Cluster
----------------------------------------

Start Kubernetes Cluster
~~~~~~~~~~~~~~~~~~~~~~~~

1. Open Vagrant shell: ``vagrant ssh``
2. Get k3s: ``curl -sfL https://get.k3s.io | sh``
3. Check nodes: ``kubectl get no``

.. warning:: To stop the Kubernetes cluster, run the command as the root user: ``shutdown -h now``


Solution 3.2 - Deploy Application to the Kubernetes Cluster
-----------------------------------------------------------

1. Run the app at a cluster: ``kubectl run POD_NAME --image=DOCKER_IMAGE_PATH``
2. Check the pod status using one of the following commands: ``kubectl get pods``
3. Deploy the app to the cluster directly from the Docker Hub: ``kubectl create deployment DEPLOYMENT_NAME --image=docker.io/DOCKERHUB_USERNAME/DOKER_IMAGE_NAME:TAG``
4. Access the application on localhost: ``kubectl port-forward POD_NAME 8080:8080``

Kubeconfig
~~~~~~~~~~

-  K3s stores the kubeconfig file under ``/etc/rancher/k3s/k3s.yaml``
-  API server - https://127.0.0.1:8080
-  Authentication mechanism - username (admin) and password

Useful kubectl Commands
~~~~~~~~~~~~~~~~~~~~~~~

-  Get the control plane and add-ons
   endpoints: ``kubectl cluster-info``
-  Get all the nodes in the cluster: ``kubectl get nodes``
-  Get extra details about the nodes: ``kubectl get nodes -o wide``
-  Get all the configuration details about the node: ``kubectl describe node NODE_NAME``
-  Get basic pods information: ``kubectl get pods`` 
-  Check the detailed information of a particular pod: ``kubectl describe pod POD_NAME`` 
-  Delete pod: ``kubectl delete pod POD_NAME``

Find out more in `kubectl Cheat Sheet <https://kubernetes.io/docs/reference/kubectl/cheatsheet/>`__

Explore kubectl predefined assets in the `video tutorial <https://www.youtube.com/watch?v=yi1kR9nDw1g>`__

.. |3.2_deploy_to_kubernetes_video| raw:: html

   <div class="video-container">
      <iframe src="https://www.youtube.com/embed/yi1kR9nDw1g" title="kubectl predefined assets" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
   </div>

|3.2_deploy_to_kubernetes_video|


Solution 3.3 - Define and Deploy Kubernetes Resources
-----------------------------------------------------

1. Create the namespace:

.. code-block:: shell 
   
   kubectl create ns demo

2. Label the namespace:

.. code-block:: shell
   
   kubectl label ns demo tier=test

3. Create the nginx-alpine deployment: 

.. code-block:: shell
   
   kubectl create deploy nginx-alpine --image=nginx:alpine  --replicas=3 --namespace demo

4. Label the deployment:

.. code-block:: shell

   kubectl label deploy nginx-alpine app=nginx tag=alpine --namespace demo

5. Expose the nginx-alpine deployment:

.. code-block:: shell

   kubectl expose deployment nginx-alpine --port=8111 --namespace demo

6. Create a config map:

.. code-block:: shell

   kubectl create configmap nginx-version --from-literal=version=alpine --namespace demo


Common Errors & How to Fix Them
-------------------------------

Permission Denied Error
~~~~~~~~~~~~~~~~~~~~~~~

The error message might look like this:

.. error:: 
   ``Unable to read /etc/rancher/k3s/k3s.yaml, please start server with --write-kubeconfig-mode to modify kube config permissions 
   error: error loading config file "/etc/rancher/k3s/k3s.yaml": open /etc/rancher/k3s/k3s.yaml: permission denied``
   
To fix **Permission Denied Error**, run one of these commands:

- ``sudo chmod 644 /etc/rancher/k3s/k3s.yaml``: changing access in k3s.yaml
- ``sudo su``: setting the superuser permissions 

Create Container Error
~~~~~~~~~~~~~~~~~~~~~~

This type of error related to the status of the pod. The error message might look like this:

.. error:: 
   ``Error: container create failed: time="2021-04-11T21:03:07Z" 
   level=error msg="container_linux.go:366: starting container process caused: 
   exec: \"cluster-kube-scheduler-operator\": executable file not found in $PATH"``

To fix **Create Container Error**, use command: ``zypper install -t pattern apparmor``