package main;

import "os";
import "fmt";
import "time";
import "regexp";
import "runtime";
import "net/http";
import "io/ioutil";
import "encoding/json";
import "github.com/sbertrang/go-htpasswd";

var version = "2.0.0";
var css = `<style>
body {
	margin: 20px;
	font-family: monospace;
}

tag {
	color: grey;
	font-style: italic;
}

tag:before {content: "("}
tag:after {content: ")"}

error {
	color: red;
}
</style>`

type Site struct {
	Port int
	Path string
	Site string

	Htpasswd string
	Authentication []string
	Authmsg string `json:"auth_msg"`

	HideFiles []string `json:"hide_files"`
	ShowFiles []string `json:"show_files"`
	AllowFiles []string `json:"allow_files"`
	BlockFiles []string `json:"block_files"`
	NoErrorPage []string `json:"no_errorpage`
	NoFilelistings []string `json:"no_filelistings"`
}

var errmsg = "\033[31m!!\033[0m "
var statusmsg = "\033[34m::\033[0m "

func handler(port string, path string, site string, config Site) http.HandlerFunc {
	fmt.Println(statusmsg + "serving \"" + site + "\" on port " + port)
	return func(w http.ResponseWriter, r *http.Request) {
		var url = r.URL.String();

		var req = path + url;
		var file, err = os.Stat(req);

		var addcss = func() {
			fmt.Fprint(w, css)
		}

		var error = func(code int, error string, msg string) {
			w.WriteHeader(code);
			for _, i := range config.NoErrorPage {
				if (regexp.MustCompile(i).MatchString(url)) {
					return;
				}
			}

			addcss();

			fmt.Fprint(w, "<b>" + error + "</b><br><br>")
			fmt.Fprint(w, "<error>" + msg + "</error>")
		}

		var notfound = func() {
			error(404, "An error occurred", "File not found!");
		}

		var authrequired = func() {
			error(401, "An error occurred", "Authorization Required!");
		}

		for _, i := range config.Authentication {
			if (regexp.MustCompile(i).MatchString(url)) {
				w.Header().Set("Authenticate", `Basic"`)
				var username, password, ok = r.BasicAuth()

				if (! ok) {
					w.Header().Add("WWW-Authenticate", `Basic realm="` + config.Authmsg + `"`)
					authrequired();
					return;
				}

				var auth, err = htpasswd.New(config.Htpasswd, htpasswd.DefaultSystems, nil)
				if (err != nil) {
					fmt.Println(errmsg + "error parsing htpasswd");
					return;
				}

				var pass = auth.Match(username, password)
				if (pass) {w.WriteHeader(200)} else {
					authrequired();
					return;
				}
			}
		}

		var serve = func() {
			if (len(config.BlockFiles) == 0 &&
				len(config.AllowFiles) == 0) {

				http.ServeFile(w, r, req);
				return;
			}

			for _, i := range config.BlockFiles {
				if (regexp.MustCompile(i).MatchString(url)) {
					notfound();
					return;
				}
			}

			if (len(config.AllowFiles) == 0) {
				http.ServeFile(w, r, req)
			} else {
				for _, i := range config.AllowFiles {
					if (regexp.MustCompile(i).MatchString(url)) {
						http.ServeFile(w, r, req);
						return;
					}
				}

				notfound();
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if (! os.IsNotExist(err)) {
			if (! file.IsDir()) {
				serve();
			} else {
				file, err = os.Stat(req + "/index.html");

				if (! os.IsNotExist(err)) {
					url = url + "/index.html"
					serve();
					return;
				}

				for _, i := range config.NoFilelistings {
					if (regexp.MustCompile(i).MatchString(url)) {
						notfound();
						return;
					}
				}

				var files, err = ioutil.ReadDir(req);
				if (err != nil) {}

				var dirsArr = []string{".."};
				var filesArr []string;

				addcss();
				fmt.Fprint(w, "<b>Directory listing for " + r.URL.String() + "</b><br><br>")

				files:
				for _, i := range files {
					for _, ii := range config.HideFiles {
						if (regexp.MustCompile(ii).MatchString(i.Name())) {
							continue files;
						}
					}
					
					if (len(config.ShowFiles) > 0) {
						for _, ii := range config.ShowFiles {
							if (! regexp.MustCompile(ii).MatchString(i.Name())) {
								continue files;
							}
						}
					}

					if (i.IsDir()) {
						dirsArr = append(dirsArr, i.Name());
					} else {
						filesArr = append(filesArr, i.Name());
					}
				}
				
				if (url == "/") {url = ""}

				for _, i := range dirsArr {
					fmt.Fprint(w, "<a href='" + url + "/" + i + "'>" + i + "/</a> <tag>Folder</tag><br>")
				}
				for _, i := range filesArr {
					fmt.Fprint(w, "<a href='" + url + "/" + i + "'>" + i + "</a> <tag>File</tag><br>")
				}
			}
		} else {
			notfound();
		}
	}
}

func server(port string, path string, site string, config Site) {
	var file, err = os.Stat(path);
	if (os.IsNotExist(err)) {
		fmt.Println(errmsg + "path doesn't exist");
		return;
	} else if (! file.IsDir()) {
		fmt.Println(errmsg + "path is not a folder");
		return;
	}

	if (port == "0") {port = "8080"}
	if (site == "") {site = "Untitled Site"}
	if (config.Authmsg == "") {config.Authmsg = "Login to view"}

	s := &http.Server{
		Addr: ":" + port,
		Handler: http.HandlerFunc(handler(port, path, site, config)),

		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}; go s.ListenAndServe();
}

func main() {
	for _, i := range os.Args[1:] {
		if (i == "-v") {
			fmt.Println("kitty: v" + version)
			fmt.Println("platform: " + runtime.GOOS + "-" + runtime.GOARCH)
			return
		} else {
			var path = regexp.MustCompile("^.*:").FindString(i);
			var port = regexp.MustCompile(":.*$").FindString(i);
			if (port == "") {port = ":8080";path = i + ":"}

			var empty Site;
			path = path[:len(path)-1];
			server(port[1:], path, path, empty)
		}
	}

	if (len(os.Args[1:]) == 0) {
		var conf = "/etc/kitty/sites.json";
		switch (runtime.GOOS) {
			case "darwin":
				conf = "/Library/Preferences/kitty/sites.json";
			case "windows":
				conf = "/kitty/sites.json";
		}

		file, err := ioutil.ReadFile(conf);
		if (err != nil) {panic(err)}

		var sites []Site;
		json.Unmarshal(file, &sites)

		for _, i := range sites {
			server(fmt.Sprintf("%d", i.Port), i.Path, i.Site, i)
		}
	}

	for true {time.Sleep(time.Second)}
}
