const fs = require("fs");
const path = require("path");
const http = require("http");

http.createServer((req, res) => {
	reqPath = path.join("site" + req.url)
	fs.readFile(reqPath, (err,data) => {
		if (err) {
			switch(err.code) {
				case "ENOENT":
					res.writeHead(404);
					res.end(JSON.stringify(err));
					break;
				case "EISDIR":
					res.writeHead(200)
					res.end("Directory listing:")
					break;
			}
			return;
		}

		res.writeHead(200);
		res.end(data);
	});
}).listen(8080);
