#!/bin/sh

GOOS=windows GOARCH=386 go build -o godiff-32.exe godiff.go
GOOS=windows GOARCH=amd64 go build -o godiff-64.exe godiff.go

