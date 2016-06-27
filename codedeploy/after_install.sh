#!/bin/bash

# make golang workspace
mkdir $HOME/go
echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> $HOME/.bashrc

# activate changes
source $HOME/.bashrc

# install golang dependency management tool
go get github.com/tools/godep

# go to golang app
cd $GOPATH/src/github.com/hoohawa/golang_web/

# restore dependencies
godep restore

# make binary
go install