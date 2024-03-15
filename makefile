# vi: set ft=makefile ts=4 sw=4

BINANRY_NAME=golang-hello-cobra

build:
	go build -o ${BINANRY_NAME} main.go

install: build
	sudo mv ${BINANRY_NAME} /usr/local/bin/

clean:
	go clean
	rm ${BINANRY_NAME}
