#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const http = require("http");
const auth = require("http-auth");

const log = require("./logs.js");
const css = fs.readFileSync(path.join(__dirname + "/main.css"), "utf8");

function server(port, site, config) {
	if (! fs.existsSync(site)) {
		log.error("path doesn't exist")
		return;
	}

	log.status(`serving "${site}" on port ${port}`)

	let server = (req, res) => {
		reqPath = path.join(site + req.url);

		let reserror = (code, error, errormsg) => {
			res.writeHead(code);
			res.write(`<style>${css}</style>`);
			res.write(`<b>${error}</b><br>`);
			res.end(`<br><error>${errormsg}</error>`);
		}

		fs.readFile(reqPath, (err, data) => {
			if (err) {
				switch(err.code) {
					case "EISDIR":
						if (config.no_filelistings) {
							if (Array.isArray(config.no_filelistings)) {

								for (let i = 0; i < config.no_filelistings.length; i++) {
									if (! new RegExp(config.no_filelistings[i]).test(reqPath)) {
										reserror(404, "An error occurred!", "File not found!");
										return;
									}
								}
							} else {
								reserror(404, "An error occurred!", "File not found!");
								return;
							}
						}

						res.writeHead(200);

						let index = path.join(`${reqPath}/index.html`)
						if (fs.existsSync(index) && fs.statSync(index).isFile()) {
							fs.readFile(index, (err, data) => {
								if (err) {throw err};
								res.end(data)
							})
							return;
						}

						res.write(`<style>${css}</style>`);
						res.write(`<b>Directory listing for ${reqPath}</b><br>`);

						let dirs = [".."];
						let files = [];

						fs.readdirSync(reqPath, (data)).forEach(i => {
							if (fs.statSync(`${reqPath}/${i}`).isDirectory()) {
								dirs[dirs.length] = i;
							} else {
								files[files.length] = i;
							}
						})

						let url = `/${req.url}/`.replace(/^\//, "")
						dirs.forEach(ii => {res.write(`<br><a href="${url}${ii}">${ii}/</a> <tag>Folder</tag>`)});
						files.forEach(ii => {res.write(`<br><a href="${url}${ii}">${ii}</a> <tag>File</tag>`)});

						res.end("");
						break;
					case "ENOENT":
						reserror(404, "An error occurred!", "File not found!");
						break;
					default:
						reserror(500, "An unhandled error occurred!", JSON.stringify(err));
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
	let config = "/etc/kitty/sites.json";
	if (process.platform == "darwin") {
		config = "/Library/Preferences/kitty/sites.json";
	} else if (process.platform == "win32") {
		config = "/kitty/sites.json";
	}

	fs.readFile(config, "utf8", (err, data) => {
		if (err) {throw err};

		config = JSON.parse(data);
		let defaultconf = {
			port: 80,
			path: "/",
			htpasswd: "",
			site: "Untitled Site",
			authentication: false,
			no_filelistings: false,
		}

		for (let i = 0; i < config.length; i++) {
			let siteconf = {...defaultconf, ...config[i]}
			new server(siteconf.port, siteconf.path, siteconf);
		}
	})
} else {
	switch(args[0]) {
		case "-v":
			log.log(`kitty: v${require("../package.json").version}`)
			log.log(`node: ${process.version}`)
			log.log(`platform: ${process.platform}-${process.arch}`)
			process.exit()
			break;
	}
	for (let i = 0; i < args.length; i++) {
		let port = parseInt(args[i].replace(/^.*:/, ""));
		let path = args[i].replace(/:.*$/, "");

		if (isNaN(port)) {
			new server(8080, path, {})
		} else {
			new server(port, path, {})
		}
	}
}

