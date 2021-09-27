# Cloud Native Foundations Program: Hands-On Course Materials, Solutions, and Project

*Learn how to structure, package, and release an application to a Kubernetes cluster, while using an automated CI/CD pipeline.*

The Cloud Native Foundations course is Phase I of the [SUSE Cloud-Native Foundations Scholarship Program](https://www.udacity.com/scholarships/suse-cloud-native-foundations-scholarship) from Udacity. The course focuses on applying software development best practices, app containerizing approaches, and releasing applications to a Kubernetes cluster using an automated CI/CD pipeline.

## About

This learning guide is an expanded version of the Cloud Native Foundations course [materials](https://github.com/udacity/nd064_course_1). The repository contains hands-on exercises and detailed step-by-step solutions developed throughout the curriculum. It also includes notes on how to correct possible errors that students may encounter while solving problems. Finally, it will include the project I'm currently working on  - Golang Service-Oriented Architecture Application.

The main goal of this work is to consolidate the knowledge gained and help students solve technical issues outside the scope of the course.

## Structure

The repository consists of three main parts: exercises, solutions, and a project.

1. All supporting materials and the [exercise_program.md]. . (exercises/exercise_program.md) can be found in the `exercises` directory, where each subfolder consists of the skeleton code and the markdown file describing the particular exercise.

2. In the directory `solutions` you will find [solutions.md](solutions/solution_list.md) file lineked to all the answers, accompanied by detailed notes on possible errors and how to fix them.

3. The `project` directory is under development.

## How to Use

- open [exercise_program.md](exercises/exercise_program.md);
- try to solve each problem from the list of exercises according to the instructions;
- if you get stuck, feel free to check the corresponding solution from the [solutions.md](solutions/solution_list.md) and reproduce it using the instructions.

## Prerequisites

- Understand the basics of HTTP (like a client, server, and internet request)
- Basic Python (data types, Functions, REST requests, web development)
- Ability to use Git, Linux machines, and Linux Command Line
- Familiar with Web application development in any language
- Familiarity with Docker and exposure to a CI/CD pipeline are not required for success in this program but are helpful skills to have.

## Course Outline

- Introduction to Cloud Native
- Architecture Considerations
- Container Orchestration
- Open Source PaaS
- Cloud Native CI/CD

## Tools, Environment & Dependencies Checklist

Make sure you have the following dependencies installed:

- [x] Python, using the instractions provided in the [official Python documentation](https://www.python.org/downloads/)
- [x] Git, using the instractions provided in the [official Git documentation](https://git-scm.com/downloads)
- [x] Docker, using the instractions provided in the [official Docker documentation](https://docs.docker.com/get-docker/)
- [x] Vagrant, using the instractions provided in the [official Vagrant documentation](https://www.vagrantup.com/downloads)
- [x] VirtualBox, using the instractions provided in the [official VirtualBox documentation](https://www.virtualbox.org/wiki/Downloads). Make sure that you have installed VirtualBox in `6.1.16` or a higher version.

## Recomendations

The current version developed and tested on [Ubuntu 20.04](https://ubuntu.com/download/desktop) using [VS Code](https://code.visualstudio.com) as IDE. While most of this stuff will work on Windows, it will take a lot of work to debug and solve some unusual problems. Use a Unix-like operating system based on the Linux kernel. This will save you a lot of time and hassle.
