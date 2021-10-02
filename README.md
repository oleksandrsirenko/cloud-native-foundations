# Cloud Native Foundations

*Learn how to structure, package, and release an application to a Kubernetes cluster using an automated CI/CD pipeline.*

## About

Cloud Native Foundations is an educational tutorial for students learning cloud-native application development.

The tutorial relates to [SUSE Cloud Native Foundations Scholarship Program](https://www.udacity.com/scholarships/suse-cloud-native-foundations-scholarship) and focuses on applying software development best practices, app containerizing approaches, and releasing applications to a Kubernetes cluster using an automated CI/CD pipeline.

The repository contains Cloud Native Foundations [hands-on exercises](https://github.com/udacity/nd064_course_1) and detailed step-by-step solutions seasoned with notes on how to correct possible errors that students encounter while solving programming challenges.

# Goal

The main goal of this project is to guide students and help them solve technical issues outside the scope of the course.

## Repository Structure

The repository consists of three main parts: exercises, solutions, and a project.

1. All supporting materials and the [exercise_program.md](exercises/exercise_program.md) can be found in the `exercises` directory, where each subfolder consists of the skeleton code and the markdown file describing the particular exercise.

2. In the directory `solutions` you will find [solutions.md](solutions/solution_list.md) file lineked to all the answers, accompanied by detailed notes on possible errors and how to fix them.

3. The `project` directory is under development.

## How to Use

- open [exercise_program.md](exercises/exercise_program.md);
- try to solve each problem from the list of exercises according to the instructions;
- if you get stuck, feel free to check the corresponding solution from the [solutions.md](solutions/solution_list.md) and reproduce it using the instructions.

## Course Outline

- Introduction to Cloud Native
- Architecture Considerations
- Container Orchestration
- Open Source PaaS
- Cloud Native CI/CD

## Required Knowledge

- Python (data types, functions, OOP).
- Basics of web development.
- Networking (REST, protocols, HTTP methods).
- Basics of Git and ability to use Git repository hosting services (GitHub, GitLab, or BitBucket).
- Shell commands in Linux.
- The knowledge of Docker, Kubernetes, and CI/CD orchestration is not required but would be helpful.

## Required Tools and Dependencies

- [x] [Python](https://www.python.org/downloads/)
- [x] [Git](https://git-scm.com/downloads)
- [x] [Docker](https://docs.docker.com/get-docker/)
- [x] [Vagrant](https://www.vagrantup.com/downloads)
- [x] [VirtualBox 6.1.16 +](https://www.virtualbox.org/wiki/Downloads)

## Recomendations

The current version developed and tested on [Ubuntu 20.04](https://ubuntu.com/download/desktop) using [VS Code](https://code.visualstudio.com) as IDE. While most of this stuff will work on Windows, it will take a lot of work to debug and solve some unusual problems. Use a Unix-like operating system based on the Linux kernel. This will save you a lot of time and hassle.
