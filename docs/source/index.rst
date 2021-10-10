.. Cloud Native Foundations Workshop documentation master file, created by
   sphinx-quickstart on Sat Oct  2 22:16:42 2021.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Welcome to Cloud Native Foundations Workshop!
====================================================

**Cloud Native Foundations Workshop** is an educational guide for students 
learning cloud-native application development.

This learning guide relates to `SUSE Cloud Native Foundations Scholarship Program 
<https://www.udacity.com/scholarships/suse-cloud-native-foundations-scholarship>`__ 
from Udacity and covers all `hands-on exercises <https://github.com/udacity/nd064_course_1>`__ 
in the Cloud Native Foundations Course. 


Objective
---------

The main goal of this project is to build a knowledge base and help solve 
technical issues outside the scope of the Cloud Native Foundation Course. 


Content
-------

The workshop contains exercises and solutions on selected topics that will help you
understand the basics of building, deploying, and maintaining cloud-native applications.  
This is NOT a comprehensive guide, but rather a focused look at a few key topics:

- Flask web development best practices.
- App containerization with Docker.
- Releasing applications to a Kubernetes cluster.
- Automation software development workflows with GitHub Actions.
- Using ArgoCD to build reliable CI/CD pipelines.

All the supporting code is in the `GitHub repository <https://github.com/oleksandrsirenko/cloud-native-foundations>`__.

Check out the :doc:`quick_start` section for further information, 
including :ref:`how to set up your workshop environment <setup_env>`.


Required Knowledge
------------------

- Python 3.
- Flask web development (introductory level).
- Networking (REST, protocols, HTTP methods).
- Git (basic commands, working with remote repositories).
- Linux shell commands.


.. toctree::
   :maxdepth: 2
   :caption: How to Use:

   quick_start
   resources

.. toctree::
   :maxdepth: 2
   :caption: Exercises:

   doc_exercises/app_endpoints_exercise
   doc_exercises/docker_exercise
   doc_exercises/kubernetes_exercise
   doc_exercises/github_actions_exercise
   doc_exercises/argocd_exercise

.. toctree::
   :maxdepth: 2
   :caption: Solutions:

   doc_solutions/app_endpoints_solution
   doc_solutions/docker_solution
   doc_solutions/kubernetes_solution
   doc_solutions/github_actions_solution
   doc_solutions/argocd_solution
