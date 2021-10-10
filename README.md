# Welcome to Cloud Native Foundations Workshop!

*Learn how to structure, package, and release an application to a Kubernetes cluster using an automated CI/CD pipeline.*

## :zzz: About

[Cloud Native Foundations Workshop]() is an educational guide for students learning cloud-native application development.

This learning guide relates to [SUSE Cloud Native Foundations Scholarship Program](https://www.udacity.com/scholarships/suse-cloud-native-foundations-scholarship)
from Udacity and covers all [hands-on exercises](https://github.com/udacity/nd064_course_1) in the Cloud Native Foundations Course.

## :dart: Objective

The main goal of this project is to build a knowledge base and help students solve technical issues outside the scope of the Cloud Native Foundation Course. 

## :card_file_box: Content

The workshop contains exercises and solutions on selected topics that will help you understand the basics of building, deploying, and maintaining cloud-native applications. This is NOT a comprehensive guide, but rather a focused look at a few key topics:

- Flask web development best practices.
- App containerization with Docker.
- Releasing applications to a Kubernetes cluster.
- Automation software development workflows with GitHub Actions.
- Using ArgoCD to build reliable CI/CD pipelines.

## :bulb: Documentation

> Find out the online version of [Cloud Native Foundations Workshop on the Read the Docs]()...comming soon.

The project documentation was developed with Sphinx and can be generated locally with a single line of code from the Linux shell.

To generate project documentation, follow these steps:

- Change the working directory to docs: `cd docs`
- Generate the documentation with: `make html`
- Open the documentation in a web browser: `firefox build/html/index.html`

## :man_student: Required Knowledge

- Python
- Flask web development (introductory level)
- Networking (REST, protocols, HTTP methods)
- Git (basic commands, working with remote repositories)
- Linux shell commands

## :computer: Required Tools

- [x] IDE of your choice ([VS Code](https://code.visualstudio.com), [Sublime](https://www.sublimetext.com/), etc.)
- [x] [Python](https://www.python.org/downloads/)
- [x] [Git](https://git-scm.com/downloads)
- [x] [Docker](https://docs.docker.com/get-docker/)
- [x] [VirtualBox 6.1.16 +](https://www.virtualbox.org/wiki/Downloads)
- [x] [Vagrant](https://www.vagrantup.com/downloads)

## :open_file_folder: Project Organization (main tree)

    │
    ├── docs               <- Workshop documetation
    ├── exercises          <- Boilerplate code for exercises
    ├── solutions          <- Complete solution code
    │
    ├── .gitignore         <- Set of patterns for files/directories to ignore
    ├── LICENSE            <- License information
    ├── README.md          <- The top-level README for developers
    ├── requirements.txt   <- The requirements file for reproducing environment
    ├── test_env.py        <- Script to test Python environment setup
    │
    └── Vagrantfile        <- Virtual machine config file (base template)

## :vulcan_salute: How to Use

The quickest way to get started is to open the [online documentation]() and walk through the tutorial. There you will find everything you need to get the job done and create your own cloud-native development solutions. 

> :pushpin: You can also download the `epub` and `pdf` version of the workshop and use it on your phone or tablet.

## :gear: Prepare to Workshop

Install VS Code, Python and Git if you have not yet. Then follow these steps to set up your development environment and get the most out of this workshop:

1. Fork this repository.
2. Clone **your fork** of this repository to your local machine.
3. Open the workshop folder from IDE (VS Code, Sublime, PyCharm).
4. Run the command  `python3 -m venv venv` from the terminal to make the virtual environment.
5. Activate the virtual environment with `source venv/bin/activate` command.
6. Change directory to documentation: `cd docs`.
7. Install documentation dependencies using `pip install -r requirements.txt` command.
8. Generate a local copy of workshop documentation by running `make html`.
9. Open the local copy of the documentation in a web browser: `firefox build/html/index.html`.