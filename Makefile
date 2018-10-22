.PHONY: run build

build:
	cd cmd/ && go build && mv cmd ../build/logger

run: build
	cd build/ && ./logger

all: run
