FROM python:latest

WORKDIR /flask
ADD flask /flask

RUN pip install -U Flask flask-socketio eventlet

EXPOSE 5000

ENTRYPOINT ["sh", "run.sh"]

