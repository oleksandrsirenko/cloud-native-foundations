
# syntax=docker/dockerfile:1

FROM python:3.8

LABEL maintainer="Oleksandr Sirenko"

COPY . /app

WORKDIR /app

RUN pip install -r requirements.txt

EXPOSE 5000

CMD [ "python", "app.py" ]