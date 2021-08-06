#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const http = require("http");
const auth = require("http-auth");

const css = {
	normal: fs.readFileSync(path.join(__dirname + "/main.css"), "utf8")
}

function server(port, site, config) {
	let server = (req, res) => {
		reqPath = path.join(site + req.url);
		fs.readFile(reqPath, (err, data) => {
			if (err) {
				switch(err.code) {
					case "EISDIR":
						res.writeHead(200);

						let index = path.join(reqPath + "index.html")
						if (fs.existsSync(index) && fs.statSync(index).isFile()) {
							fs.readFile(index, (err, data) => {
								res.end(data)
							})
							return;
						}

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

						dirs.forEach(ii => {res.write(`<br><a href="${ii}">${ii}/</a> <tag>Folder</tag>`)});
						files.forEach(ii => {res.write(`<br><a href="${ii}">${ii}</a> <tag>File</tag>`)});

						res.end("");
						break;
					case "ENOENT":
						res.writeHead(404);
						res.write(`<style>${css.normal}</style>`);
						res.write("<b>An error occurred!</b><br>");
						res.end("<br><err>File not found!</err>");
						break;
					default:
						res.writeHead(404);
						res.write(`<style>${css.normal}</style>`);
						res.write("<b>An unhandled error occurred!</b><br>");
						res.end(`<br><err>${JSON.stringify(err)}</err>`);
				}
				return;
			}

			res.writeHead(200);
			res.end(data);
		});
	}

	if (config.authentication) {
		let basic = auth.basic({
			file: config.htpasswd
		})

		http.createServer(basic.check((req, res) => {
			server(req, res);
		}), (req, res) => {
			server(req, res);
		}).listen(port);
	} else {
		http.createServer((req, res) => {
			server(req, res);
		}).listen(port);
	}
}

args = process.argv.splice(2, process.argv.length)

if (args[0] == undefined) {
	fs.readFile("/etc/kitty/sites.json", "utf8", (err, data) => {
		if (err) {throw err};

		config = JSON.parse(data);

		for (let i = 0; i < config.length; i++) {
			new server(config[i].port, config[i].path, config[i]);
		}
	})
} else {
	if (args.length == 1) {
		new server(8080, args[0], {})
	}
}

