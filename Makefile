BIN = /usr/bin/kitty

deps:
	@cd src;go install

compile: deps
	@mkdir -p build
	@rm build/* -rf
	cd src;GOOS=linux   go build -o ../build/kitty-linux main.go 
	cd src;GOOS=darwin  go build -o ../build/kitty-macos main.go 
	cd src;GOOS=windows go build -o ../build/kitty.exe   main.go 

install: compile
	@cp build/kitty-linux $(BIN)
	@chmod 755 $(BIN)

start: deps
	@cd src;go run main.go ../examples
