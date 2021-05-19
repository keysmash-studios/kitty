<p align="center">
    <img width="200px" src="assets/paw.png"></img><br>
    <b>Kitty</b><br>
    <a href="https://github.com/keysmash-studios/kitty/projects/1">Todo</a> - <a href="https://github.com/keysmash-studios/kitty/issues?q=is%3Aissue+is%3Aopen">Open Issues</a>
<p>

This is a simple webserver, it does allow you to host other services, that is, you can't host wordpress on it.

However, you can host static web technologies, and overall make a static HTTP server.

<br><br><br>

Installing
----------

Currently the only way to run kitty is directly with Node, that is you clone the repo, install the dependencies (Node, NPM Modules, GNU make (optional)) an example below:

```sh
$ git clone https://github.com/keysmash-studios/kitty

$ cd kitty

$ npm i # or npm install

$ make start # or node src/main.js
```

In the future the plan is to both provide the current method and also a singular executable, that is it'll have Node and the modules built in, so you just have the main executable `kitty`

However it is not the current plan, as currently that's not needed as it's nowhere near where you can use it properly anyway.

<br><br><br>

History of the name
-------------------

It's quite simple funny enough, when "learning" about webservers in IT class we were being told to use WampServer to host a Wordpress site.

I then questioned "Why don't you just use caddy", I then got told "Kitty, make our own, and name it kitty"

(Khaokci edit here, I was the smartass who said it)

Thanks Caddy for the inspiration to our name.
