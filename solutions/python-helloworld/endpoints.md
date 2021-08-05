# Exercise 1 Endpoints for Application Status

Extend the Python Flask application with `/status` and `/metrics` endpoints, considering the following requirements:

- Both endpoints should return an HTTP 200 status code
- Both endpoints should return a JSON response e.g. `{"user": "admin"}`. (Note: the JSON response can be hardcoded at this stage)
- The `/status` endpoint should return a response similar to this example: `result: OK - healthy`
- The `/metrics` endpoint should return a response similar to this example: `data: {UserCount: 140, UserCountActive: 23}`

Feel free to check the [solution video](https://www.youtube.com/watch?v=Kj_hGnViybg) where detailed operations are demonstrated.

## Solution 1: Endpoints

Add to route for each logic like:

```python
@app.route("/status")
def status():
    response = app.response_class(
        response=json.dumps({"result": "OK - healthy"}),
        status=200,
        mimetype="application/json"
    )

    return response
```

```python
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
```

See complete code in the `app.py`
