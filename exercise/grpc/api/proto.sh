protoc --proto_path=$GOPATH/src:. --go_out=plugins=grpc:. --go_opt=paths=source_relative api.proto
