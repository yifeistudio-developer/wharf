#!/bin/zsh
protoc --go_out=../golang/alps --go_opt=paths=source_relative --go-grpc_out=../golang/alps --go-grpc_opt=paths=source_relative alps.proto
