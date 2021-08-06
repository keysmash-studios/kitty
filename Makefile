BIN = /usr/bin
SRC = $(BIN)/kitty-src

install:
	@mkdir $(SRC) -p
	@cp src/* $(SRC)
	@cp package.json $(SRC)
	@cp src/start.sh $(BIN)/kitty
	@chmod 755 $(BIN)/kitty $(SRC)/start.sh $(SRC)/main.js
	@cd $(SRC);npm i $(SRC)

uninstall:
	@rm $(BIN)/kitty $(BIN)/kitty-src

entr:
	@ls "$(PWD)"/src/* | entr -r make -s start

start:
	@node src/main.js
