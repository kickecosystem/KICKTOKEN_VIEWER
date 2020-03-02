all: build run

build:
	go build

run:
	./kicktoken_viewer -l ws://128.0.0.1:8546