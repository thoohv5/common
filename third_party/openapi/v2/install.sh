#/bin/bash

protoc --proto_path=. --proto_path=../../../third_party --go_out=paths=source_relative:. ./annotations.proto
protoc --proto_path=. --proto_path=../../../third_party --go_out=paths=source_relative:. ./openapi.proto