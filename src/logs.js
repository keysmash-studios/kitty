const colors = require("colors")

function log(string, type) {
	switch(type) {
		case "error":
			console.log("!! ".red + string)
			break;
		case "status":
			console.log(":: ".blue + string)
			break;
		case "success":
			console.log("<> ".green + string)
			break;
		default:
			console.log(string)
	}
}

var error = (string) => {log(string, "error")}
var status = (string) => {log(string, "status")}
var success = (string) => {log(string, "success")}

module.exports = {
	log,
	error,
	status,
	success
}
