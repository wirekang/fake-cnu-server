#!/bin/sh

GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/fake_cnu_server_windows.exe main.go
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/fake_cnu_server_linux main.go
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/fake_cnu_server_mac_amd main.go
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o bin/fake_cnu_server_mac_arm main.go
