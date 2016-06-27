#!/bin/bash

#update any packages if needed
apt-get update -y
apt-get upgrade -y

apt-get install git -y

# install golang
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz