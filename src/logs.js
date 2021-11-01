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

module.exports = {
	log,
	error: (str) => {log(str, "error")},
	status: (str) => {log(str, "status")},
	success: (str) => {log(str, "success")}
}
