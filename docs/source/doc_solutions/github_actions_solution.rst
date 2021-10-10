4. GitHub Actions
=================

.. warning::
  To securely access the Docker Hub repository, you need to create two secrets in your GitHub 
  account: ``DOCKER_HUB_USERNAME`` and ``DOCKER_HUB_ACCESS_TOKEN`` respectively. Read more `here <https://docs.docker.com/ci-cd/github-actions/>`__.


Solution 4.1 Dockerize Application with GitHub Actions
------------------------------------------------------

To Dockerize your application using GitHub Actions follow these steps:

1. Open the ``docker-build.yaml`` file under the ``solution/`` directory.
2. Edit this file to meet your needs (correct app name, version, etc.).
3. Save changes you have made.
4. On GitHub, open your fork of the workshop repository.
5. Put ``docker-build.yaml`` to the ``.github/workflows`` directory to execute.

The ``docker-build.yaml`` for Python Flask application should look like this:

.. code-block:: yaml

    # This is a basic workflow to help you get started with Actions

    name: Docker Build and Push

    # Controls when the action will run. Triggers the workflow on push or pull request
    # events but only for the master branch
    on:
      push:
        branches: [ master ]
      pull_request:
        branches: [ master ]

    # A workflow run is made up of one or more jobs that can run sequentially or in parallel
    jobs:
      # This workflow contains a single job called "build"
      build:
        # The type of runner that the job will run on
        runs-on: ubuntu-latest

        # Steps represent a sequence of tasks that will be executed as part of the job
        steps:
          -
            name: Check Out Repo
            uses: actions/checkout@v2
          -
            name: Set up QEMU
            uses: docker/setup-qemu-action@v1
          -
            name: Login to DockerHub
            uses: docker/login-action@v1 
            with:
              username: ${{ secrets.DOCKER_HUB_USERNAME }}
              password: ${{ secrets.DOCKER_HUB_ACCESS_TOKEN }}
          -
            name: Set up Docker Buildx
            uses: docker/setup-buildx-action@v1
          -
            name: Build and push
            uses: docker/build-push-action@v2
            with:
              context: ./
              file: ./Dockerfile
              platforms: linux/amd64
              push: true
              tags: ${{ secrets.DOCKER_HUB_USERNAME }}/python-helloworld:latest


Watch the `solution <https://www.youtube.com/watch?v=rx3N8MImP9k>`__ video lesson.

.. |4.1_github_actions_solution| raw:: html

  <div class="video-container">
    <iframe src="https://www.youtube.com/embed/rx3N8MImP9k" title="Dockerize the app with GitHub Actions" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
  </div>

|4.1_github_actions_solution|