.DEFAULT_GOAL := build

clean:
	sudo GOPATH=$$GOPATH go clean -r
	sudo rm -f /usr/local/bin/msg

build:
	go build -o ./bin/msg

install:
	sudo cp ./bin/msg /usr/local/bin
	sudo chmod +x /usr/local/bin/msg
