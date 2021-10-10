1. Application Endpoints and Logging
====================================

.. note:: 
   Make sure you have :ref:`prepared the environment <lesson_1_prep>` for this task.


Solution 1.1 - Extend Application Endpoints
-------------------------------------------

To extend Python Flask application with ``/status`` and ``/metrics``
endpoints, follow these steps:  


1. Open ``app.py`` file in ``exercises/python-helloworld`` folder.
2. Register routes ``"/status"`` and ``"/metrics"`` to the app route.
3. Add the logic to those routes, according to the examples below:

.. code-block:: python
    :emphasize-lines: 1

    @app.route("/status")
    def status():
        response = app.response_class(
            response=json.dumps({"result": "OK - healthy"}),
            status=200,
            mimetype="application/json"
        )

        return response

.. code-block:: python
    :emphasize-lines: 1

    @app.route("/metrics")
    def metrics():
        response = app.response_class(
            response=json.dumps(
                {
                    "status": "success",
                    "code": 0,
                    "data": {"UserCount": 140, "UserCountActive": 23}
                }
            ),
            status=200,
            mimetype="application/json"
        )

        return response

Watch the `video <https://www.youtube.com/watch?v=Kj_hGnViybg>`__ to understand the API Endpoints solution in more detail.

.. |1.1_endpoints_video| raw:: html

    <div class="video-container">
      <iframe src="https://www.youtube.com/embed/Kj_hGnViybg" title="Application Endpoints" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    </div>

|1.1_endpoints_video|


Solution 1.2 - Implement Application Logging
--------------------------------------------

For log messages Flask uses standard Python ``logging`` library. In Flask it implemented within ``app.logger`` method, 
which can be used in the same way and allow you to handle your own masseges. 
To add basic logging functionality to your Flask application follow the next steps:

1. Open ``app.py`` file in ``exercises/python-helloworld`` folder.
2. Import ``logging`` library.
3. Add log collection logic for each route you want to track:

.. code-block:: python
    :emphasize-lines: 4

    @app.route("/metrics")
    def metrics():
        ... 
        app.logger.info("Metrics request successfull.")


- Add logging configuration to main function to save the logs in a file:

.. code-block:: python
    :emphasize-lines: 2

    if __name__ == "__main__":
        logging.basicConfig(
            filename="app.log",
            level=logging.DEBUG,
            format="%(levelname)s:%(asctime)s:%(name)s:%(message)s",
        )
        
        app.run(host="0.0.0.0")


Watch the `video <https://www.youtube.com/watch?v=rdoXsSx1ghk>`__ to understand the logging basics in more detail.

.. |1.2_logging_video| raw:: html

    <div class="video-container">
      <iframe src="https://www.youtube.com/embed/rdoXsSx1ghk" title="Application Logging" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen>></iframe>
    </div>

|1.2_logging_video|


.. note:: 
    In this tutorial, we discover the very basics of logging. The complete solution code is in `solutions/python-helloworld/app.py <https://github.com/oleksandrsirenko/cloud-native-foundations/blob/main/solutions/python-helloworld/app.py>`__.
    To improve the app logging facility and create a more sophisticated solution explore the additional materials for the exercise.

