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
