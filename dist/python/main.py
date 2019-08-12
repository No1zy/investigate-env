from urllib.parse import urlparse

url = 'http://exampel.com/ok'

u = urlparse(url)

try:
    print('Scheme: ' + u.scheme)
    print('netloc: ' + u.netloc)
    print('path: ' + u.path)
    print('params: ' + u.params)
    print('query: ' + u.query)
    print('fragment: ' + u.fragment)
    print('username: ' + str(u.username))
    print('password: ' + str(u.password))
    print('hostname: ' + str(u.hostname))
    print('port: ' + str(u.port))
except:
    import traceback
    print(traceback.print_exc())
