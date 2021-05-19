const fs = require("fs");
const path = require("path");
const http = require("http");

const css = {
	normal: fs.readFileSync(path.join(__dirname + "/main.css"), "utf8")
}

http.createServer((req, res) => {
	reqPath = path.join("site" + req.url);
	fs.readFile(reqPath, (err,data) => {
		if (err) {
			switch(err.code) {
				case "EISDIR":
					res.writeHead(200);
					res.write(`<style>${css.normal}</style>`);
					res.write(`<b>Directory listing for ${reqPath}</b><br>`);

					let dirs = [];
					let files = [];

					fs.readdirSync(reqPath, (data)).forEach(i => {
						if (fs.statSync(reqPath + i).isDirectory()) {
							dirs[dirs.length] = i;
						} else {
							files[files.length] = i;
						}

					})

					dirs.forEach(ii => {
						res.write(`<br><a href="${ii}">${ii}/</a>`);
					})

					files.forEach(ii => {
						res.write(`<br><a href="${ii}">${ii}</a>`);
					})

					res.end("");
					break;
				default:
					res.writeHead(404);
					res.write(`<style>${css.normal}</style>`);
					res.end(JSON.stringify(err));
			}
			return;
		}

		res.writeHead(200);
		res.end(data);
	});
}).listen(8080);
