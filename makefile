.DEFAULT_GOAL := default

clean:
	sudo GOPATH=$$GOPATH go clean -r
	sudo rm -f /usr/local/bin/msg

default:
	go build -o ./bin/msg

install:
	sudo cp ./bin/msg /usr/local/bin
	sudo chmod +x /usr/local/bin/msg
