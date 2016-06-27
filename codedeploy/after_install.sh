#!/bin/bash

# install dependencies
pwd >> /home/ubuntu/log.txt
echo $PATH >> /home/ubuntu/log.txt
go get github.com/gin-gonic/gin

# go to golang app
cd /home/ubuntu/go/src/github.com/hoohawa/golang_web/

# make binary
go install