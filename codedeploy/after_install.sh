#!/bin/bash

# install dependencies
whoami >> /home/ubuntu/log.txt
pwd >> /home/ubuntu/log.txt
echo $PATH >> /home/ubuntu/log.txt
/usr/local/bin/go get github.com/gin-gonic/gin

# go to golang app
cd /home/ubuntu/go/src/github.com/hoohawa/golang_web/

# make binary
/usr/local/bin/go install