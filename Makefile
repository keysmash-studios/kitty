BIN = /usr/bin
SRC = $(BIN)/kitty-src

install:
	@npm i
	@mkdir $(SRC) -p
	@cp src/* $(SRC)
	@cp package.json $(SRC)
	@cp scripts/start.sh $(BIN)/kitty
	@chmod 755 $(BIN)/kitty $(SRC)/index.js
	@cd $(SRC);npm i $(SRC)

uninstall:
	@rm $(BIN)/kitty $(BIN)/kitty-src -rf

compile:
	@rm build -rf
	@mkdir build
	@npm i
	@cp scripts/install.sh build
	@node_modules/.bin/pkg .

entr:
	@ls "$(PWD)"/src/* | entr -r make -s start

start:
	@node src/index.js site
