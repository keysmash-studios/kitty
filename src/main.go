package main;

import "os";
import "fmt";
import "time";
import "regexp";
import "runtime";
import "net/http";
import "io/ioutil";
import "encoding/json";

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

	HideFiles []string `json:"hide_files"`
	ShowFiles []string `json:"show_files"`
	AllowFiles []string `json:"allow_files"`
	BlockFiles []string `json:"block_files"`
	NoErrorPage []string `json:"no_errorpage`
	NoFilelistings []string `json:"no_filelistings"`
}

func handler(port string, path string, site string, config Site) http.HandlerFunc {
	fmt.Println(":: serving \"" + site + "\" on port " + port)
	return func(w http.ResponseWriter, r *http.Request) {
		var url = r.URL.String();

		var req = path + url;
		var file, err = os.Stat(req);

		var addcss = func() {
			fmt.Fprint(w, css)
		}

		var error = func(code uint, error string, msg string) {
			w.WriteHeader(404);
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
		file, err := ioutil.ReadFile("/etc/kitty/sites.json");
		if (err != nil) {panic(err)}

		var sites []Site;
		json.Unmarshal(file, &sites)

		for _, i := range sites {
			server(fmt.Sprintf("%d", i.Port), i.Path, i.Site, i)
		}
	}

	for true {time.Sleep(time.Second)}
}
