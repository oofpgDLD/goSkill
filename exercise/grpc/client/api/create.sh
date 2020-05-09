#!/bin/sh

#_path=$(go list -f '{{.Dir}}' "gitlab.33.cn/chat/chat33/service/user/")
protoc --go_out=plugins=grpc:. *.proto
