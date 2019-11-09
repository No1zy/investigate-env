from flask import Flask, render_template
from flask_socketio import SocketIO, emit, send

import os

app = Flask(__name__)
app.config['SECRET_KEY'] = os.getenv('SECRET_KEY', default="secret!")
socketio = SocketIO(app, cors_allowed_origins="*")

@app.route('/')
def index():
    return render_template("index.html")

@app.route('/flag')
def flag():
    return render_template("flag")

@socketio.on('connect', namespace='/socket.io/')
def connect():
    app.logger.info("connected")
    send('connected')

@socketio.on('disconnect', namespace='/socket.io/')
def disconnect():
    print('Client disconnected')

@socketio.on('message', namespace='/socket.io/')
def hello(message):
    app.logger.info("recived message:" + message)
    send('Hello, world!')

if __name__ == '__main__':
    socketio.run(app, host='0.0.0.0', debug=True)
    
