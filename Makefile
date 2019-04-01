all: build

build:
	- cd /chaincode/src/uniblock && go build
	- cd /web && go build