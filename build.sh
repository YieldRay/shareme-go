#!/usr/bin/env bash

bin="shareme.exe"

if [ "$1" == "c" ]; then
    go build -ldflags "-s -w" -o $bin
    upx --best $bin
fi

if [ "$1" == "b" ]; then
    go build -ldflags "-s -w" -o $bin
    upx --brute $bin
fi

if [ "$1" == "d" ]; then
    go run main.go && rm main.exe
fi

if [ "$1" == "" ]; then
    go build -o $bin
fi
