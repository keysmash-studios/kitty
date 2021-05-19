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
			res.write(`<style>${css.normal}</style>`);
			switch(err.code) {
				case "EISDIR":
					res.writeHead(200);
					res.write(`<b>Directory listing for ${reqPath}</b><br>`);
					fs.readdirSync(reqPath, (data)).forEach(i => {
						res.write(`<br><a href="${i}">${i}</a>`);
					})
					res.end("");
					break;
				default:
					res.writeHead(404);
					res.end(JSON.stringify(err));
			}
			return;
		}

		res.writeHead(200);
		res.end(data);
	});
}).listen(8080);
