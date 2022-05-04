package handlers_test

import (
	"proto2http/handlers"
	"proto2http/target"
	"testing"

	"github.com/emicklei/proto"
)

func TestMessageHandler(t *testing.T) {
	var d = &handlers.Dependency{
		Log:     noplog.Logger,
		Verbose: true,
		Output:  &target.Proto{},
	}

	var msg = &proto.Message{
		Name: "SampleMessage",
		Comment: &proto.Comment{
			Lines: []string{"Message comment"},
		},
		Elements: []proto.Visitee{
			&proto.NormalField{
				Field: &proto.Field{
					Name:     "Field1",
					Type:     "string",
					Sequence: 0,
					Comment: &proto.Comment{
						Lines: []string{"Field comment"},
					},
				},
				Repeated: true,
				Optional: true,
				Required: false,
			},
			&proto.MapField{
				Field: &proto.Field{
					Name:     "Field2",
					Type:     "map<string, string>",
					Sequence: 1,
				},
				KeyType: "string",
			},
		},
	}

	d.MessageHandler(msg)

	if d.Output.Messages[0].Name != "SampleMessage" {
		t.Errorf("Expected message name to be 'SampleMessage', got '%s'", d.Output.Messages[0].Name)
	}

	if d.Output.Messages[0].Comment != "Message comment" {
		t.Errorf("Expected message comment to be 'Message comment', got '%s'", d.Output.Messages[0].Comment)
	}

	if len(d.Output.Messages[0].Fields) != 2 {
		t.Errorf("Expected 2 message field, got %d", len(d.Output.Messages[0].Fields))
	}

	if d.Output.Messages[0].Fields[0].Name != "Field1" {
		t.Errorf("Expected field name to be 'Field1', got '%s'", d.Output.Messages[0].Fields[0].Name)
	}

	if d.Output.Messages[0].Fields[0].Type != "string" {
		t.Errorf("Expected field type to be 'string', got '%s'", d.Output.Messages[0].Fields[0].Type)
	}

	if d.Output.Messages[0].Fields[0].Repeated != true {
		t.Errorf("Expected field repeated to be true, got false")
	}

	if d.Output.Messages[0].Fields[0].Optional != true {
		t.Errorf("Expected field optional to be true, got false")
	}

	if d.Output.Messages[0].Fields[0].Required != false {
		t.Errorf("Expected field required to be false, got true")
	}

	if d.Output.Messages[0].Fields[0].Comment != "Field comment" {
		t.Errorf("Expected field comment to be 'Field comment', got '%s'", d.Output.Messages[0].Fields[0].Comment)
	}

	if d.Output.Messages[0].Fields[1].Name != "Field2" {
		t.Errorf("Expected field name to be 'Field2', got '%s'", d.Output.Messages[0].Fields[1].Name)
	}

	if d.Output.Messages[0].Fields[1].Type != "map<string, string>" {
		t.Errorf("Expected field type to be 'map<string, string>', got '%s'", d.Output.Messages[0].Fields[1].Type)
	}
}

func TestMessageHandlerRpcIntegration(t *testing.T) {
	var d = &handlers.Dependency{
		Log:     noplog.Logger,
		Verbose: true,
		Output:  &target.Proto{},
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
				RequestType: "SampleRequest",
				ReturnsType: "SampleResponse",
			},
			&proto.RPC{
				Name: "TwoRPC",
				Comment: &proto.Comment{
					Lines: []string{"Another RPC comment"},
				},
				RequestType: "EmptyRequest",
				ReturnsType: "OkResponse",
			},
		},
	}

	d.ServiceHandler(srv)

	var msgs = []*proto.Message{
		{
			Name: "SampleRequest",
			Comment: &proto.Comment{
				Lines: []string{"Message comment"},
			},
			Elements: []proto.Visitee{
				&proto.NormalField{
					Field: &proto.Field{
						Name:     "Field1",
						Type:     "string",
						Sequence: 0,
						Comment: &proto.Comment{
							Lines: []string{"Field comment"},
						},
					},
					Repeated: true,
					Optional: true,
					Required: false,
				},
				&proto.MapField{
					Field: &proto.Field{
						Name:     "Field2",
						Type:     "map<string, string>",
						Sequence: 1,
					},
					KeyType: "string",
				},
			},
		},
		{
			Name: "SampleResponse",
			Elements: []proto.Visitee{
				&proto.NormalField{
					Field: &proto.Field{
						Name:     "User",
						Type:     "string",
						Sequence: 0,
					},
					Repeated: false,
					Optional: false,
					Required: false,
				},
			},
		},
	}

	for _, msg := range msgs {
		d.MessageHandler(msg)
	}

	if len(d.Output.Services[0].RPCs) != 2 {
		t.Errorf("Expected 2 RPCs, instead got %d", len(d.Output.Services[0].RPCs))
	}

	if d.Output.Services[0].RPCs[0].Name != "OneRPC" {
		t.Errorf("Expected first RPC name to be OneRPC, instead got %s", d.Output.Services[0].RPCs[0].Name)
	}

	if d.Output.Services[0].RPCs[0].Request.Name != "SampleRequest" {
		t.Errorf("Expected first RPC's request name to be 'SampleRequest', instead got %s", d.Output.Services[0].RPCs[0].Request.Name)
	}

	if len(d.Output.Services[0].RPCs[0].Request.Fields) != 2 {
		t.Errorf("Expected first RPC's request fields to be 2, instead got %d", len(d.Output.Services[0].RPCs[0].Request.Fields))
	}

}
