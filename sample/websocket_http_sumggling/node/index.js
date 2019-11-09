const express = require("express")
const fs = require('fs')

let app = new express()

app.get("/", (req, res) => {
    fs.readFile(__dirname + "/index.html", 'utf8', (err, text) => {
        res.send(text)
    })
})

app.listen(3000, () => {console.log("app listen port: 3000")})

