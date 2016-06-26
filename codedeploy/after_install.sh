#!/bin/bash

# make golang workspace
cd $HOME
mkdir go
echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
echo "export PATH=$PATH:$HOME/bin" >> $HOME/.bashrc

# activate changes
source $HOME/.bashrc

# install dependencies
go get github.com/gin-gonic/gin
go get github.com/golang/protobuf/proto
go get github.com/manucorporat/sse
go get golang.org/x/net/context

# go to golang app
cd $GOPATH/src/github.com/hoohawa/golang_web/

# make binary
go install