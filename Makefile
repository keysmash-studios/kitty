BIN = /usr/bin
SRC = $(BIN)/kitty-src

install:
	@npm i
	@mkdir $(SRC) -p
	@cp src/* $(SRC)
	@cp package.json $(SRC)
	@cp src/start.sh $(BIN)/kitty
	@chmod 755 $(BIN)/kitty $(SRC)/start.sh $(SRC)/index.js
	@cd $(SRC);npm i $(SRC)

uninstall:
	@rm $(BIN)/kitty $(BIN)/kitty-src -rf

compile:
	@rm build -rf
	@mkdir build
	@npm i
	@node_modules/.bin/pkg .

entr:
	@ls "$(PWD)"/src/* | entr -r make -s start

start:
	@node src/index.js
