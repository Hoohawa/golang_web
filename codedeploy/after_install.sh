#!/bin/bash

# install dependencies
go get github.com/gin-gonic/gin
go get github.com/golang/protobuf/proto
go get github.com/manucorporat/sse
go get golang.org/x/net/context

# go to golang app
cd /home/ubuntu/go/src/github.com/hoohawa/golang_web/

# make binary
go install