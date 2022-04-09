<p align="center">
    <img width="200px" src="assets/paw.png"><br>
    <img width="400px" src="assets/code.png">
<p>

Kitty is a simple webserver, it allows you to easily host static websites, with minimal configuration and ease of use.

<br><br><br>

Installing
----------

### Development

You can either run it directly in the repo like so:

```sh
$ git clone https://github.com/keysmash-studios/kitty

$ cd kitty

$ make start
```

### Executable

You can just download a build from the [Releases page](https://github.com/keysmash-studios/kitty/releases) for a more stable build.
Or compile the upstream version:

```sh
$ git clone https://github.com/keysmash-studios/kitty

$ cd kitty

$ make compile

# then run your platform's executable
$ build/kitty-[linux|macos|.exe]
# preferably rename them or something
```

You can also simply run `make install` if you're on Linux.

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

Thanks Caddy for the inspiration to our name.
