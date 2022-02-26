#!/usr/bin/env bash
VERSION=0.1.0
go build -o gredir/gredir ./main.go
cp config.yaml gredir/
cp install.sh gredir/
tar cvf gredir-$VERSION.tar gredir/
gzip gredir-$VERSION.tar
rm -rf gredir
