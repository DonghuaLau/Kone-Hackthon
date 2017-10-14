
export GOPATH=$(shell pwd)
export GOBIN=${GOPATH}/bin
export GOSRC=${GOPATH}/src

all:
	@echo "Usage: make <target>"

webapp:
	@cd ${GOSRC}/$@; \
	go install; \
	cd -;

clean:
	rm -rf bin/*
	rm -rf pkg/*
