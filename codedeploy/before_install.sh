#!/bin/bash

#update any packages if needed
apt-get update -y
apt-get upgrade -y

apt-get install git -y

# install golang
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz

# make golang workspace
echo "export GOPATH=/home/ubuntu/go" >> /home/ubuntu/.bashrc
echo "export PATH=$PATH:/home/ubuntu/go/bin" >> /home/ubuntu/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/ubuntu/.bashrc

# activate changes
source /home/ubuntu/.bashrc