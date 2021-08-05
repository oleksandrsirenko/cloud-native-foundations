# Exercise 1.2 Application Logging

Logging is a core factor in increasing the visibility and transparency of an application. When in troubleshooting or debugging scenarios, it is paramount to pin-point the functionality that impacted the service. This exercise will focus on bringing the logging capabilities to an application.

At this stage, you have extended the Hello World application to handle different endpoints. Once an endpoint is reached, a log line should be recorded showcasing this operation.

In this exercise, you need to further develop the Hello World application collect logs, with the following requirements:

- A log line should be recorded the timestamp and the requested endpoint e.g. `"{{TIMESTAMP}}, {{ ENDPOINT_NAME}} endpoint was reached"`
- The logs should be stored in a file with the name `app.log`. Refer to the [logging Python module](https://docs.python.org/3/library/logging.html#logging.basicConfig) for more details.
- Enable the collection of Python logs at the `DEBUG` level. Refer to the [logging Python module](https://docs.python.org/3/library/logging.html#logging.basicConfig) for more details.

Note: For the environment setup, follow the instructions in the previous exercise.

Feel free to check the [solution video](https://www.youtube.com/watch?v=rdoXsSx1ghk) where detailed operations are demonstrated.

## Solution 1.2 Logging

Add log collection logic to the appropriate functions for each route you want to track, like:

```python
@app.route("/metrics")
def metrics():
    ... 
    app.logger.info("Metrics request successfull")
```

Also add the logic to save the logs in a file to main function (need to import `logging` library first):

```python
logging.basicConfig(filename='app.log',level=logging.DEBUG)
```

See complete code in the `app.py`
