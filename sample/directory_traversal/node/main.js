const fs = require('fs')
const path = require('path');
const express = require('express')

let app = express()

const base = '/data/'

app.get('/', (req, res) => {

    if (req.query.file === '') {
        res.send('require file param')
        return
    }

    const file = req.query.file
    console.log(req.query.file)

    if (file.indexOf('../') === -1) {
        fs.readFile(path.join(base,decodeURIComponent(file)), (err, data) => {
            if (err) console.log(err)
    
            res.send(data)
        })
    } else {
        res.send('Detect: "../" ')
    }
})

app.listen('3000')
