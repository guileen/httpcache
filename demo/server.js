var http = require('http')

http.createServer(function(req, res) {
    console.log(req.method, req.url, req.headers)
    res.end("method:"+req.method + req.url)
}).listen(3000)
