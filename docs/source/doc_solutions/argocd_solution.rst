5. Continuous Delivery Fundamentals
===================================

.. note:: 
  The preparation steps are the same as in the exercise ``3. Kubernetes Cluster``
  Make sure you have :ref:`prepared the environment <lesson_3_prep>`


Solution 5.1 - Deploy the Application Using ArgoCD
--------------------------------------------------

Quick Start with ArgoCD
~~~~~~~~~~~~~~~~~~~~~~~

1. Open Vagrant shell using command:

.. code-block:: shell
  
  vagrant ssh

2. Create namespace:

.. code-block:: shell
  
  kubectl create namespace argocd

3. Install ArgoCD:

.. code-block:: shell
  
  kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

3. Check ArgoCD pods:

.. code-block:: shell
  
  kubectl get po -n argocd

4. Get all application services:

.. code-block:: shell
  
  kubectl get svc -n argocd


Create and Apply YAML Manifests
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. Warning::
  Do not forget to change ``GITHUB_USERNAME`` inside the ``YAML`` files below to your GitHub user name!

argocd-server-nodeport.yaml
"""""""""""""""""""""""""""

1. Create a file within the vim shell: ``vim argocd-server-nodeport.yaml``
2. Write the manifest:

.. code-block:: yaml

  apiVersion: v1
  kind: Service
  metadata:
    annotations:
    labels:
      app.kubernetes.io/component: server
      app.kubernetes.io/name: argocd-server
      app.kubernetes.io/part-of: argocd
    name: argocd-server-nodeport
    namespace: argocd
  spec:
    ports:
    - name: http
      port: 80
      protocol: TCP
      targetPort: 8080
      nodePort: 30007
    - name: https
      port: 443
      protocol: TCP
      targetPort: 8080
      nodePort: 30008
    selector:
      app.kubernetes.io/name: argocd-server
    sessionAffinity: None
    type: NodePort

3. Escape the vim shell: ``esc``
4. Save changes: ``:wq``
5. Check the manifest: ``cat argocd-server-nodeport.yaml``
6. Apply the manifest: ``kubectl apply -f argocd-server-nodeport.yaml``

nginx-alpine.yaml
"""""""""""""""""

1. Create a file within the vim shell: ``vim nginx-alpine.yaml``
2. Write the manifest:

.. code-block:: yaml
  
  apiVersion: argoproj.io/v1alpha1
  kind: Application
  metadata:
    name: nginx-alpine
    namespace: argocd
  spec:
    destination:
      namespace: default
      server: https://kubernetes.default.svc
    project: default
    source:
      path: solutions/kubernetes/manifests 
      repoURL: https://github.com/GITHUB_USERNAME/cloud-native-foundations
      targetRevision: HEAD
    # Sync policy
    syncPolicy: {}

3. Escape the vim shell: ``esc``
4. Save changes: ``:wq``
5. Check if everything was written correctly: ``cat nginx-alpine.yaml``
6. Apply created yaml manifest:: ``kubectl apply -f nginx-alpine.yaml``

After preparing all the required manifests, check the application using the command: 

.. code-block:: shell

  kubectl get application -n argocd

Use ArgoCD in your browser at
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

-  http://192.168.50.4.30007
-  http://192.168.50.4.30008


Solutions 5.2 - Create a Helm Chart with ArgoCD
-----------------------------------------------

The Helm chart is defined in the ``Chart.yaml`` file, which contains the
API version, name and version of the chart:

.. code-block:: yaml

  apiVersion: v1
  name: nginx-deployment
  description: Install Nginx deployment manifests 
  keywords:
    - nginx 
  version: 1.0.0
  maintainers:
    - name: GITHUB_USERNAME


Solution 5.3 - Create a YAML Values Files
-----------------------------------------

An example of the ``values.yaml`` file:

.. code-block:: yaml

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

The above configuration represents the default parameters of application
deployment if it is not overwritten by a different values file.

Below is an example of the ``values-prod.yaml`` file, which will
override the default parameters:

.. code-block:: yaml

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


Solution 5.4 - Create ArgoCD Applications
-----------------------------------------

ArgoCD application CRD for the ``nginx-prod.yaml`` deployment:

.. code-block:: yaml

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
        repoURL: https://github.com/GITHUB_USERNAME/cloud-native-foundations/tree/main/solutions/argocd
        targetRevision: HEAD

.. note:: 
  The ``nginx-staging.yaml``, ``values-staging.yaml``, and ``nginx-prod.yaml`` files can be found 
  in the project repository `solutions/helm/nginx-deployment <https://github.com/oleksandrsirenko/cloud-native-foundations/tree/main/solutions/helm/nginx-deployment>`__


Common Errors & How to Fix Them
-------------------------------

Error Validating Data
~~~~~~~~~~~~~~~~~~~~~

Pay attention when copying/pasting to build manifests! The letter ``a`` 
on the first line of a created manifest is usually lost after the file is saved. 
This error turns ``apiVersion`` into ``piVersion`` and raises the following error:

.. error:: 
  ``error: error validating "argocd-server-nodeport.yaml": error validating data:
  apiVersion not set; if you choose to ignore these errors, turn validation off with --validate=false``

To fix **Error Validating Data**, open the yaml file you working on and correct the typo 
by turning ``piVersion`` back to ``apiVersion``.