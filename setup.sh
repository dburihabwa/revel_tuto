#! /bin/bash
ROOT=~/
# ROOT=/gfs/$(whoami)/
echo "mkdir ~/gocode"
mkdir $ROOT/gocode
touch ~/.gorc
GOPATH=$ROOT/gocode
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
# Download code
cd $GOPATH/src
git clone https://github.com/dburihabwa/revel_tuto going
