5. Continuous Delivery Fundamentals
===================================

.. |argocd_demo| raw:: html

  <div class="video-container">
    <iframe src="https://www.youtube.com/embed/TZgLkCFQ2tk" title="Argo CD Demo" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
  </div>

|argocd_demo|

In this exercise, you will use Argo CD to automate the delivery of an
application to a Kubernetes cluster.

Continuous Delivery (CD) is the ability to get code changes reliably to
production environments. This practice should be automated and should
enable developers to provide value to consumers efficiently.


Preparation
-----------

Prepare your environment by following :ref:`Preparation <lesson_3_prep>` 
section in the exercise :doc:`3. Kubernetes Cluster <kubernetes_exercise>`


Exercise 5.1 - Deploy the Application Using Argo CD
---------------------------------------------------

1. Install Argo CD by referring to the `installation guide <https://argoproj.github.io/argo-cd/getting_started/>`__.
2. Expose the ``argocd-server`` using a ``NodePort`` service on port ``30007`` on HTTP and ``30008`` on HTTPS. The YAML manifest for the NodePost service can be found in the tutorial repository.
3. Access the Argo CD UI by going to ``https://192.168.50.4.30008`` or ``https://192.168.50.4.30007``.
4. Insert login credentials by using the credentials guide.
5. Create an Argo CD application named ``nginx-alpine`` and use the manifests provided in the course repository.


Configuration Managers
----------------------

For the management of multiple declarative Kubernetes manifests, a
templating layer is necessary, especially if the application is
replicated across different regions. For this purpose configuration
managers, such as Helm and Kustomize, were introduced.

This exercise will focus on creating your first Helm chart to deploy
multiple Nginx applications using the same template and multiple input
files.

Exercise 5.2 - Create a Helm Chart with Argo CD
-----------------------------------------------

Using the manifests provided in the course repository, create a helm
chart (Chart.yaml, templates, values.yaml) that will template the
following parameters:

- **namespace name**
- **replica count**
- image:
   - name
   - tag
   - pull policy
- resources:
   - requests for CPU and memory
- service:
   - port
   - type (e.g. ClusterIP)
- **configmap data** (e.g. the key-value pair)

The chart details should be as following:

- **name:** nginx-deployment
- **version:** 1.0.0
- **keywords:** nginx

Exercise 5.3 - Create a YAML Values Files
-----------------------------------------

Once the Helm chart is available make sure that a default
``values.yaml`` file is available. This values file will be used as a
default input file for the Helm chart. The ``values.yaml`` file should
have the following specification:

``values.yaml``
  -  namespace name: demo
  -  replica count: 3
  -  image repository: nginx
  -  image tag: alpine
  -  image pull policy: IfNotPresent
  -  resources: CPU 50m and memory 256Mi
  -  service type: ClusterIP
  -  service port: 8111
  -  configmap data: "version: alpine"

Next, create two values files with the following specifications:

``values-staging.yaml``
  -  namespace name: staging
  -  replica count: 1
  -  image repository: nginx
  -  image tag: 1.18.0
  -  resources: CPU 50m and memory 128Mi
  -  configmap data: "version: 1.18.0"

``values-prod.yaml``
  -  namespace name: prod
  -  replica count: 2
  -  image repository: nginx
  -  image tag: 1.17.0
  -  resources: CPU 70m and memory 256Mi
  -  service port: 80
  -  configmap data: "version: 1.17.0"

Exercise 5.4 - Create Argo CD Applications
------------------------------------------

Using the values files above (values-prod, values-staging),
create two Argo CD applications, ``nginx-staging`` and ``nginx-prod``
respectively. These should deploy the nginx Helm Chart referencing each
input values files.

Additional Resources
--------------------

1. `Getting Started with Argo CD <https://argo-cd.readthedocs.io/en/stable/getting_started/>`__
2. `Argo CD Installation Video <https://www.youtube.com/watch?v=TJrSM31Jj_8>`__
3. `Guide To GitOps <https://www.weave.works/technologies/gitops/>`__
4. `Helm Quickstart Guide <https://helm.sh/docs/intro/quickstart/>`__
5. `Helm Deployment with Argo CD Video Tutorial <https://www.youtube.com/watch?v=VyuVFtp2-2M&t=26s>`__
6. `CI/CD Guides for DevOps Engineers - 8 Video Playlist <https://youtube.com/playlist?list=PLHq1uqvAteVsSsrnZimHEf7NJ1MlRhQUj>`__
