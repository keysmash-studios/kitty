<p align="center">
    <img width="200px" src="assets/paw.png"><br>
    <img width="400px" src="assets/code.png">
<p>

This file/page is dedicated completely to go over the various options
and ways to configure kitty, all it's options, all it's features and
everything alike.

Command Line
------------

On the command line you can start kitty, and multiple sites with very
little hassle, simply as the command line arguments add the path to a
folder, it'll then host it. To host multiple sites you'll obviously have
to specify the port, as you can't host two sites on the same port, you
can do that by adding `:<port>` after the path.

Lets say we have a 2 folders, "site" and "site2", we want to host them
on port 80, 8080 and 8081 we simply do:

```sh
$ kitty site:80 site site2:8081
```

Since it defaults to 8080 not providing one will use 8080, you can also
host the same folder on multiple ports.

More command line options will be added soon!

Config File
-----------

The file locations and or names may be changed in the future, more
importantly I'm not sure if the Windows location will stick, as I simply
placed it there because I don't understand the file locations for
Windows, hence that was just the only place I could think of putting it,
make a pull request if there's a better location for it.

The file does not get created by kitty and has to be manually setup.

Linux: `/etc/kitty/sites.json` <br>
Windows: `C:/kitty/sites.json` <br>
macOS: `/Library/Preferences/kitty/sites.json`

An example of the config file is inside the `examples/` folder.

### Options

`sites.json` as the name suggests only houses options for sites, it is a
big array with an object (a site) which then has options inside of it.

#### port

`Default: 80`

This is straight forward, it is simply the port which the site will be
running on, whether it be 8080, 80, or a completely different port.
Generally you can only use a port above 0, and below 65536, however this
is a restriction by the OS itself (from what I know) and could change.
You also can't use anything below or equal to 1024 without root/admin
permissions, again a restriction by the OS.

#### site

`Default: "Untitled Site"`

This setting is what's printed when mutliple sites are started to make it
easier to identify each individual site more clearly, it is also planned to be
used for more than just that in the future.

#### path

`Default: "/"`

Again very straight forward, the path for the site, if a path that
doens't exist is provided it'll error out.

#### authentication

`Default: false`

`authentication` enables the `htpasswd` option, the `htpasswd` option
takes in a path to a `htpasswd` file, similar to that of Nginx or alike,
you can generate one with various online tools or use Apache's tools or
simply anything alike it.

#### htpasswd

`Default: ""`

When a valid path is provided and it's enabled it'll prompt the user to
enter credentials when accessing the site. It's up to the browser how
long it'll save the login, and I don't know of any way I can prolong or
edit it, not even tell the browser to listen to my recommendation of how
long it should stay logged in.

#### no_filelistings

`Default: false`

This takes either an array or boolean, if set to `true` it'll disable
all file listings and give a 404 error instead, if set to `false` it'll
do nothing and show the file listings. If set to an array, when
accessing a page it'll check if the path is matched with a regEx
pattern, that is, the array houses regEx patterns and when matched it'll
disable file listings.