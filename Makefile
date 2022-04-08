BIN = /usr/bin/kitty

compile:
	@mkdir -p build
	@rm build/* -rf
	GOOS=linux   go build -o build/kitty-linux src/main.go 
	GOOS=darwin  go build -o build/kitty-macos src/main.go 
	GOOS=windows go build -o build/kitty.exe   src/main.go 

install: compile
	@cp build/kitty-linux $(BIN)
	@chmod 755 $(BIN)

start:
	@go run src/main.go examples
