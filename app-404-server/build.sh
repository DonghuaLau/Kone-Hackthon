#!/bin/sh

GOPATH=$(pwd)
GOBIN=${GOPATH}/bin
GOSRC=${GOPATH}/src

export GOPATH=${GOPATH}
export GOBIN=${GOBIN}
export GOSRC=${GOSRC}


cd ${GOSRC}/webapp
go install
cd -;

#cp ${GOSRC}/webapp/views/*.tmpl  bin/views/ -rf
