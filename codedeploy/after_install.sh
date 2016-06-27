#!/bin/bash

# make golang workspace
mkdir $HOME/go
export GOPATH=$HOME/go >> $HOME/.bashrc
export PATH=$PATH:$GOPATH/bin >> $HOME/.bashrc
export PATH=$PATH:/usr/local/go/bin >> $HOME/.bashrc

pwd >> log.txt
cd $HOME
pwd >> log.txt
whoami >> log.txt
echo $GOPATH >> log.txt
echo $PATH >> log.txt

# install golang dependencies
go get github.com/gin-gonic/gin
go get golang.org/x/net/websocket

# go to golang app
cd $GOPATH/src/github.com/hoohawa/golang_web/

# make binary
go install