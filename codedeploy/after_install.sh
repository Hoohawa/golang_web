#!/bin/bash

# make golang workspace
mkdir $HOME/go
echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> $HOME/.bashrc

# activate changes
source $HOME/.bashrc

# install dependencies
echo $GOPATH >> /home/ubuntu/log.txt
whoami >> /home/ubuntu/log.txt
pwd >> /home/ubuntu/log.txt
echo $PATH >> /home/ubuntu/log.txt

go get github.com/gin-gonic/gin

# go to golang app
cd /home/ubuntu/go/src/github.com/hoohawa/golang_web/

# make binary
go install