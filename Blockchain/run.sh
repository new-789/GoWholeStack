#!/bin/bash
rm blockChain
rm *.db
go build -o blockChain *.go
./blockChain
