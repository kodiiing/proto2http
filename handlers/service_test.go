package handlers_test

import (
	"testing"

	"github.com/kodiiing/proto2http/handlers"
	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

func TestServiceHandler(t *testing.T) {
	var d = &handlers.Dependency{
		Output: &target.Proto{},
	}

	var srv = &proto.Service{
		Name: "SampleService",
		Comment: &proto.Comment{
			Lines: []string{"Service comment"},
		},
		Elements: []proto.Visitee{
			&proto.RPC{
				Name: "OneRPC",
				Comment: &proto.Comment{
					Lines: []string{"RPC comment"},
				},
			},
			&proto.RPC{
				Name: "TwoRPC",
				Comment: &proto.Comment{
					Lines: []string{"Another RPC comment"},
				},
			},
		},
	}

	d.ServiceHandler(srv)

	if len(d.Output.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(d.Output.Services))
	}

	if d.Output.Services[0].Name != "SampleService" {
		t.Errorf("Expected service name to be 'SampleService', got '%s'", d.Output.Services[0].Name)
	}

	if d.Output.Services[0].Comment != "Service comment" {
		t.Errorf("Expected service comment to be 'Service comment', got '%s'", d.Output.Services[0].Comment)
	}

	for _, rpc := range d.Output.Services[0].RPCs {
		if rpc.Name != "OneRPC" && rpc.Name != "TwoRPC" {
			t.Errorf("Expected RPC name to be OneRPC or TwoRPC, got %s", rpc.Name)
		}

		if rpc.Comment == "" {
			t.Errorf("Expected RPC comment to be set, got empty string")
		}
	}
}
