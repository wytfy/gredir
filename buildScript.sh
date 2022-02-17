#!/usr/bin/env bash
mkdir build
go build -o build/gredir ./main.go
cp config.yaml build/
tar cvf gredir.tar build/
gzip gredir.tar