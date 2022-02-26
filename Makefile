BIN = /usr/bin
SRC = $(BIN)/kitty-src

install:
	@npm i
	@rm -rf $(BIN)/kitty $(SRC)
	@mkdir $(SRC)/src -p
	@cp src/* $(SRC)/src
	@cp package.json $(SRC)
	@cp scripts/start.sh $(BIN)/kitty
	@chmod 755 $(BIN)/kitty $(SRC)/src/index.js
	@cd $(SRC);npm i $(SRC)

uninstall:
	@rm $(BIN)/kitty $(BIN)/kitty-src -rf

compile:
	@rm build -rf
	@mkdir build
	@npm i
	@cp scripts/install.sh build
	@node_modules/.bin/pkg . -t node16-linux,node16-macos,node16-win

entr:
	@ls "$(PWD)"/src/* | entr -r make -s start

start:
	@node src/index.js examples/site
