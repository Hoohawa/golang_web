#!/bin/bash

# make golang workspace
echo "export GOPATH=/home/ubuntu/go" >> /home/ubuntu/.bashrc
echo "export PATH=$PATH:/home/ubuntu/go/bin" >> /home/ubuntu/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/ubuntu/.bashrc

# activate changes
source /home/ubuntu/.bashrc

# install dependencies
whoami >> /home/ubuntu/log.txt
pwd >> /home/ubuntu/log.txt
echo $PATH >> /home/ubuntu/log.txt

go get github.com/gin-gonic/gin

# go to golang app
cd /home/ubuntu/go/src/github.com/hoohawa/golang_web/

# make binary
go install