#!/bin/bash

#update any packages if needed
apt-get update -y
apt-get upgrade -y

# install golang
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz

# make golang workspace
cd $HOME/go
echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
echo "export PATH=$PATH:$HOME/go/bin" >> $HOME/.bashrc

# activate changes
source $HOME/.bashrc