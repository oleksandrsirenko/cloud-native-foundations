3. Kubernetes Cluster
=====================

.. |3.0_kubernetes_intro| raw:: html

  <div class="video-container">
    <iframe src="https://www.youtube.com/embed/VnvRFRk_51k" title="GitHub Actions - Supercharge your GitHub Flow" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
  </div>

|3.0_kubernetes_intro|

In the previous exercise, you put the Go app into a Docker container and pushed 
it to the Docker Hub repository. So now you need to deploy this application from 
your Docker Hub repository to your local Kubernetes cluster.

.. note:: 
    This exercise is slightly different from the original one - it was modified 
    according to more practical use cases and connect student progress from the 
    previous lessons.

.. _lesson_3_prep:

Preparation
-----------

Install VirtualBox and Vagrant
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

1. Install VirtualBox: ``sudo apt install virtualbox``
2. Install Vagrant: ``sudo apt-get update && sudo apt-get install vagrant``

Up Virtual Machine
~~~~~~~~~~~~~~~~~~

1. Change directory to exercises: ``cd exercises``
2. Init virtual machine: ``vagrant init``
3. Up the virtual machine: ``vagrant up``
4. Check the status of virtual machine: ``vagrant status``

Useful Vagrant Commands
~~~~~~~~~~~~~~~~~~~~~~~

-  Shut down virtual machine forcefully: ``vagrant halt``
-  Suspend the virtual machine: ``vagrant suspend``
-  Restart virtual machine: ``vagrant up``

Find out more in `Vagrant Cheat Sheet <https://linuxacademy.com/site-content/uploads/2017/12/vagrant-cheatsheet-Linux-Academy.pdf>`__


Exercise 3.1 - Create Kubernetes Cluster
----------------------------------------

- Create k3s Kubernetes cluster on the local machine using Virtual Box and Vagrant.

Exercise 3.2 - Deploy Application to the Kubernetes Cluster
-----------------------------------------------------------

- Run go-helloworld app at the Kunernetes cluster from the Docker Hub.
- Deploy go-helloworld to the Kubernetes cluster from the Docker Hub.

Exercise 3.3 - Define and Deploy Kubernetes Resources
-----------------------------------------------------

- Deploy the following resources using the kubectl command:

    - a namespace:
        - name: ``demo``
        - label: ``tier: test``
    - a deployment:
        - image: ``nginx:alpine``
        - name: ``nginx-apline``
        - namespace: ``demo``
        - replicas: 3
        - labels: ``app: nginx``, ``tag: alpine``
    - a service:
        - expose the above deployment on port ``8111``
        - namespace: ``demo``
    - a configmap:
        - name: ``nginx-version``
        - containing key-value pair: ``version=alpine``
        - namespace: ``demo``

.. note:: 
    Nginx is one of the public Docker images, that you can access 
    and use for your exercises or testing purposes.

Make sure the following tasks are completed:

.. |check| raw:: html

    <input checked=""  type="checkbox">

.. |uncheck| raw:: html

    <input type="checkbox">

.. |br| raw:: html

    <br/>

|check| You have created a Namespace |br|
|check| You have created a Deployment |br|
|check| You have created a Service |br|
|check| You have created a Configmap |br|


Additional Resourses
--------------------

1. `Kubernetes Tutorial for Beginners - 4 Hour Video Course <https://youtu.be/X48VuDVv0do>`__
2. `Official Kubernetes Tutorials <https://kubernetes.io/docs/tutorials/>`__
3. `Kubernetes Hands-on Labs <https://katacoda.com/courses/kubernetes/>`__
4. `Vagrant Documentation Resources <https://www.vagrantup.com/docs>`__
5. `Vagrant Cheat Sheet <https://linuxacademy.com/site-content/uploads/2017/12/vagrant-cheatsheet-Linux-Academy.pdf>`__
6. `K3s Lightweight Kubernetes <https://k3s.io/>`__
7. `Stopping and starting Kubernetes cluster <https://www.ibm.com/docs/en/fci/1.0.3?topic=kubernetes-stopping-starting-cluster>`__
8. `Organizing Cluster Access Using kubeconfig Files <https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/>`__
9. `How to Manage Kubernetes With Kubectl <https://rancher.com/learning-paths/how-to-manage-kubernetes-with-kubectl/>`__
10. `Explore kubectl Cheat Sheet <https://kubernetes.io/docs/reference/kubectl/cheatsheet/>`__
11. `Kubernetes Config file <https://community.suse.com/posts/cluster-this-is-your-admin-do-you-read>`__
12. `Using kubectl to Create a Deployment <https://kubernetes.io/docs/tutorials/kubernetes-basics/deploy-app/deploy-intro/>`__
13. `How to Delete Pods from a Kubernetes Node <https://www.bluematador.com/blog/safely-removing-pods-from-a-kubernetes-node>`__
14. `Use Port Forwarding to Access Applications in a Cluster <https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/>`__