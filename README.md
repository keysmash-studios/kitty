<p align="center">
    <img width="200px" src="assets/paw.png"><br>
    <img width="400px" src="assets/code.png">
<p>

This is a simple webserver, it does allow you to host other services, that is, you can't host wordpress on it.

However, you can host static web technologies, and overall make a static HTTP server.

<br><br><br>

Installing
----------

### Development

You can either run it directly in the repo like so:

```sh
$ git clone https://github.com/keysmash-studios/kitty

$ cd kitty

$ npm i # or npm install

$ make start # or node src/index.js
```

### NodeJS Dependant

```sh
$ git clone https://github.com/keysmash-studios/kitty

$ cd kitty

$ make install

# to remove: (on Linux and maybe macOS)
$ make uninstall
```

### Executable

You can just download a build from the [Releases page](https://github.com/keysmash-studios/kitty/releases) for a more stable build.
Or compile the upstream version:

```sh
$ git clone https://github.com/keysmash-studios/kitty

$ cd kitty

$ make compile

# then run your platform's executable
$ build/kitty-[linux|macos|win.exe]
# preferably rename them or something

# or just run the install script for Linux+macOS
$ cd build; ./install.sh
```

`make install` simply copies `src/start.sh` to `/usr/bin/kitty` copies the content of `src/` to `/usr/bin/kitty-src/` and the `package*.json` files as well, then installs the needed npm modules. This technically means we're vendoring the dependencies, in the sense that it installs the modules separate from system modules.

You can also if you have [`entr`](http://eradman.com/entrproject/) installed use that to restart kitty when changes are made to the source code, this is very useful for developing/testing kitty.

```sh
$ make entr
# It'll then run kitty
# And on source changes restart kitty
```

<br><br><br>

Kitty is...
-----------------------

 * Not a replacement to nginx, apache, or alike
 * Not a full on replacement for any web server

 * Supposed to make it easy to host static files
 * Supposed to make it easy to host multiple sites

<br><br><br>

Tracking Progress
-----------------

All progress, bugs, issues and alike is tracked through [git-bug](https://github.com/MichaelMure/git-bug), with it installed, simply create and identity (for first time use) and use the webui or termui, or cli if you're into that.

```
$ git bug user create

$ git bug webui
# Or
$ git bug termui
```

However feel free to use the GitHub issues as well for reporting bugs, as anybody can do that without credentials.

<br><br><br>

History of the name
-------------------

It's quite simple funny enough, when "learning" about webservers in IT class we were being told to use WampServer to host a Wordpress site.

I then questioned "Why don't you just use caddy", I then got told "Kitty, make our own, and name it kitty"

(Khaokci edit here, I was the smartass who said it)

Thanks Caddy for the inspiration to our name.
