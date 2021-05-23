#!/bin/bash
rm blockChain
rm *.db
rm *.dat
go build -o blockChain *.go
./blockChain
