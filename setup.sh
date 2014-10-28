#! /bin/bash

echo "mkdir ~/gocode"
mkdir ~/gocode
touch ~/.gorc
GOPATH=~/gocode
export GOPATH=$GOPATH
echo "export GOPATH=$GOPATH" > ~/.gorc

echo "go get github.com/revel/revel"
go get github.com/revel/revel
echo "go get github.com/revel/cmd/revel"
go get github.com/revel/cmd/revel

export PATH=$PATH:$GOPATH/bin
echo "export PATH=$PATH" >> ~/.gorc
echo "" >> ~/.bashrc
echo "# GO CODE" >> ~/.bashrc
echo "source ~/.gorc" >> ~/.bashrc
