package main

import "testing"

func TestExecute(t *testing.T) {
	ret := (&configWrapper{
		Name:       "Config",
		Comment:    "//配置",
		HasComment: true,
		Methods: []*method{
			{
				Name:       "Server",
				RetType:    "string",
				Comment:    "//服务配置",
				HasComment: true,
			},
			{
				Name:       "Server1",
				RetType:    "string1",
				Comment:    "//服务配置1",
				HasComment: true,
			},
		},
	}).execute()

	t.Logf(ret)
}
