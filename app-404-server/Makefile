
export GOPATH=$(shell pwd)
export GOBIN=${GOPATH}/bin
export GOSRC=${GOPATH}/src

all:
	@echo "Usage: make <target>"

hello:
	@cd ${GOPATH}/src/$@; \
	go install; \
	cd -;

main:
	@cd ${GOPATH}/src/$@; \
	go install; \
	cd -;

webapp:
	@cd ${GOSRC}/$@; \
	go install; \
	cd -;

clean:
	rm -rf bin/*
	rm -rf pkg/*
