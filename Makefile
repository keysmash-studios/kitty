tests:
	@echo "No tests are implemented"

entr:
	@ls $PWD/src/* | entr make start

start:
	@node src/main.js
