1. Application Endpoints and Logging
====================================

.. |1.0_flask_cs50| raw:: html

  <div class="video-container">
    <iframe src="https://www.youtube.com/embed/x_c8pTW8ZUc" title="GitHub Actions - Supercharge your GitHub Flow" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
  </div>

|1.0_flask_cs50|

In this exercise, you will use a minimal Flask application that looks like this:

.. code-block:: python

   from flask import Flask

   app = Flask(__name__)

   @app.route("/")
   def hello():
      return "<p>Hello World!</>"

   if __name__ == "__main__":
      app.run(host="0.0.0.0")

Your tasks are to extend application endpoints and implement the logging functionality.
If you are not familiar with Flask, start with the `Flask Quickstart documentation <https://flask.palletsprojects.com/en/2.0.x/quickstart/>`__ and
watch the `Flask CS50 video tutorial <https://youtu.be/x_c8pTW8ZUc>`__ at the top of this page.


.. _lesson_1_prep:

Preparation
-----------

.. note:: 
   Make sure you have :ref:`installed all essentials (IDE, Python, Git) <setup_env>`   

Follow these steps to set up the environment for the exercise:

1. Create the virtual environment: ``python3 -m venv venv``
2. Activate the virtual environment: ``source venv/bin/activate``
3. Change the working directory: ``cd exercises/python-helloworld``
4. Install dependencies: ``pip install -r requirements.txt``
5. Run the application: ``python app.py``
6. Test the application in a web browser at: http://192.168.1.3:5000/


Exercise 1.1 - Extend Application Endpoints
-------------------------------------------

The endpoint of a web application is simply a URL where a web browser can access 
the application. So, to extend application endpoints, you need to define new URL 
routes and create view functions to match those routes. For a better understanding, 
carefully read two sections of the Flask documentation:

- `URL Route Registrations <https://flask.palletsprojects.com/en/2.0.x/api/#url-route-registrations>`__
- `Response Objects <https://flask.palletsprojects.com/en/2.0.x/api/#response-objects>`__

.. Tip::
   Use ``@app.route()`` decorator to bind a view function for the application endpoint.

**Extend your Python Flask app with** ``/status`` **and** ``/metrics`` **endpoints, considering the following requirements:**

-  Both endpoints must return an status code ``HTTP 200``
-  Both endpoints must return a JSON response e.g. ``{"user": "admin"}``
-  The ``/status`` endpoint should return a response similar to this
   example: ``result: OK - healthy``
-  The ``/metrics`` endpoint should return a response similar to this example: 
   ``data: {UserCount: 140, UserCountActive: 23}``

Exercise 1.2 - Implement Application Logging
--------------------------------------------

Flask uses standard Python `logging <https://docs.python.org/3/library/logging.html#module-logging>`__ module. 
Messages about your Flask application are logged with `app.logger <https://flask.palletsprojects.com/en/2.0.x/api/#flask.Flask.logger>`__, which takes the 
same name as app.name. This logger can also be used to log your own messages [3].

.. tip::
   Configure logging for your application as soon as possible when the program starts.

**Add a log collection logic for each endpoint in your Flask application according 
to the following requirements:**

- Logs should be kept in a separate file named ``app.log``.
- You need to collect logs at the ``DEBUG`` level.
- Each log must be in the same format and include information about the logging level, actual time, endpoint name, and message.


Additional Resources
--------------------

1. `Flask Quick Start <https://flask.palletsprojects.com/en/2.0.x/quickstart/>`__
2. `Flask URL Route Registrations <https://flask.palletsprojects.com/en/2.0.x/api/#url-route-registrations>`__
3. `Flask Logging Documentaion <https://flask.palletsprojects.com/en/2.0.x/logging/?highlight=logging>`__
4. `Logging Facility for Python <https://docs.python.org/3/library/logging.html>`__
5. `Python Basic Logging Tutorial <https://docs.python.org/3/howto/logging.html#logging-basic-tutorial>`__
6. `Python Logging Cookbook <https://docs.python.org/3/howto/logging-cookbook.html#logging-cookbook>`__
7. `Python Advanced Logging Tutorial <https://docs.python.org/3/howto/logging.html#logging-advanced-tutorial>`__
8. `CS50 2020 - Lecture 9 - Flask <https://youtu.be/x_c8pTW8ZUc>`__
9. `Flask Application Video Tutoral from Tech with Tim <https://youtu.be/mqhxxeeTbu0>`__
10. `Learn Flask for Python - freeCodeCamp Video Tutorial <https://youtu.be/Z1RJmh_OqeA>`__