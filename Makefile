tests:
	@echo "No tests are implemented"

entr:
	@ls "$(PWD)"/src/* | entr -r make -s start

start:
	@node src/main.js
