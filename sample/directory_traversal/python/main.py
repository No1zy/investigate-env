import os

PREFIX = 'data/'

path = '{{ .VARIABLE }}'

data_path = os.path.join(PREFIX, path)
full_path = os.path.abspath(data_path)

data = ''
with open(full_path, 'r') as f:
    data = f.read()

print(data)
