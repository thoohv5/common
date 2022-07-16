package main

import (
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestConfig(t *testing.T) {
	gen, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		Parameter: proto.String("Mdir/conf.proto=/Users/thooh/Sites/github.com/thoohv5/common/cmd/protoc-gen-go-configs"),
		ProtoFile: []*descriptorpb.FileDescriptorProto{
			{
				Name:    proto.String("conf.proto"),
				Syntax:  proto.String(protoreflect.Proto3.String()),
				Package: proto.String("ipam_plus.api"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("git.zdns.cn/lirui/zdns/app/service/ipam_plus-service/internal/conf;conf"),
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("missing go_package option: New(req) = nil, want error")
	}

	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range gen.Files {
		// if !f.Generate {
		// 	continue
		// }
		generateFile(gen, f)
	}
}
