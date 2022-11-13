const http = require('http'),
      fs = require('fs'),
      url = require('url'),
      path = require('path'),
      mimeModel = require('./model/getMime.js');

http.createServer(function (request, response){
    var pathname = url.parse(request.url).pathname;
    var mime = mimeModel.getMime(path.extname(pathname));

    console.log("Request for " + pathname + " received.");
    fs.readFile(pathname.substr(1), function (err, data){
        if(err){
            console.log(err);
            response.writeHead(404, {'Content-Type': ''+mime+'; charset="utf-8"'});
            response.end("404 NOt FOUND");
        }
        else{
            response.writeHead(200, {'Content-Type': ''+mime+'; charset="utf-8"'});
            response.write(data.toString());
        }
        response.end();
    });
}).listen(8080);

console.log('Server running at http://127.0.0.1:8080/');
