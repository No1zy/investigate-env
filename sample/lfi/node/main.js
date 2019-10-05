const fs = require('fs')

url = '{{ .VARIABLE }}'

const params = new URL(url).searchParams
base = '/data/'
path = params.get('file')

if (path.indexOf('..') === -1) {
    fs.readFile(base + path, (err, data) => {
        if (err) throw err

        console.log(data)
    })
} else {
    console.log('Detect: ".." ')
}
