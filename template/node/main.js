const url = require('url')

const u = '{{ .VARIABLE }}'

const parsed_url1 = url.parse(u)

console.log(parsed_url1)

const parsed_url2 = new URL(u)

console.log(parsed_url2)
