package handlers_test

import (
	"errors"
	"os"
	"proto2http/handlers"
	"proto2http/target"
	"testing"

	"github.com/emicklei/proto"
)

func TestFixtures_Auth(t *testing.T) {
	var d = &handlers.Dependency{
		Verbose: true,
		Log:     noplog.Logger,
		Output:  &target.Proto{},
	}

	file, err := os.Open("./fixtures/auth.proto")
	if err != nil {
		t.Fatalf("Failed to open proto file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			t.Errorf("error during closing file: %v", err)
		}
	}()

	protoFile := proto.NewParser(file)

	parsedProtoFile, err := protoFile.Parse()
	if err != nil {
		t.Fatalf("error parsing proto file: %v", err)
	}

	proto.Walk(
		parsedProtoFile,
		proto.WithPackage(d.PackageHandler),
		proto.WithService(d.ServiceHandler),
		proto.WithMessage(d.MessageHandler),
		proto.WithEnum(d.EnumHandler),
	)

	if d.Output.Name != "grpc.testing" {
		t.Errorf("expected name to be grpc.testing, got %s", d.Output.Name)
	}

	if len(d.Output.Services) != 1 {
		t.Errorf("expected 1 service, got %d", len(d.Output.Services))
	}

	if d.Output.Services[0].Name != "TestService" {
		t.Errorf("expected service name to be TestService, got %s", d.Output.Services[0].Name)
	}

	if len(d.Output.Services[0].RPCs) != 1 {
		t.Errorf("expected 1 RPCs, got %d", len(d.Output.Services[0].RPCs))
	}
}

func TestFixtures_HelloWorld(t *testing.T) {
	var d = &handlers.Dependency{
		Verbose: true,
		Log:     noplog.Logger,
		Output:  &target.Proto{},
	}

	file, err := os.Open("./fixtures/helloworld.proto")
	if err != nil {
		t.Fatalf("Failed to open proto file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			t.Errorf("error during closing file: %v", err)
		}
	}()

	protoFile := proto.NewParser(file)

	parsedProtoFile, err := protoFile.Parse()
	if err != nil {
		t.Fatalf("error parsing proto file: %v", err)
	}

	proto.Walk(
		parsedProtoFile,
		proto.WithPackage(d.PackageHandler),
		proto.WithService(d.ServiceHandler),
		proto.WithMessage(d.MessageHandler),
		proto.WithEnum(d.EnumHandler),
	)

	if d.Output.Name != "helloworld" {
		t.Errorf("expected name to be helloworld, got %s", d.Output.Name)
	}

	if len(d.Output.Services) != 1 {
		t.Errorf("expected 1 service, got %d", len(d.Output.Services))
	}

	if d.Output.Services[0].Name != "Greeter" {
		t.Errorf("expected service name to be Greeter, got %s", d.Output.Services[0].Name)
	}

	if len(d.Output.Services[0].RPCs) != 1 {
		t.Errorf("expected 1 RPCs, got %d", len(d.Output.Services[0].RPCs))
	}

	var rpc = d.Output.Services[0].RPCs[0]

	if rpc.Name != "SayHello" {
		t.Errorf("expected RPC name to be SayHello, got %s", rpc.Name)
	}

	if rpc.Comment != " Sends a greeting" {
		t.Errorf("expected RPC comment to be Sends a greeting, got %s", rpc.Comment)
	}

	if rpc.Request.Name != "HelloRequest" {
		t.Errorf("expected RPC request name to be HelloRequest, got %s", rpc.Request.Name)
	}

	if len(rpc.Request.Fields) != 1 {
		t.Errorf("expected 1 fields on RPC fields, got %d instead", len(rpc.Request.Fields))
	}

	if rpc.Request.Fields[0].Name != "name" {
		t.Errorf("expected RPC request field name to be name, got %s", rpc.Request.Fields[0].Name)
	}

	if rpc.Request.Fields[0].Type != "string" {
		t.Errorf("expected RPC request field type to be string, got %s", rpc.Request.Fields[0].Type)
	}

	if rpc.Response.Name != "HelloReply" {
		t.Errorf("expected RPC response name to be HelloReply, got %s", rpc.Response.Name)
	}

	if len(rpc.Response.Fields) != 1 {
		t.Errorf("expected 1 fields on RPC response, got %d instead", len(rpc.Response.Fields))
	}

	if rpc.Response.Fields[0].Name != "message" {
		t.Errorf("expected RPC response field name to be message, got %s", rpc.Response.Fields[0].Name)
	}

	if rpc.Response.Fields[0].Type != "string" {
		t.Errorf("expected RPC response field type to be string, got %s", rpc.Response.Fields[0].Type)
	}
}

func TestFixtures_RouteGuide(t *testing.T) {
	var d = &handlers.Dependency{
		Verbose: true,
		Log:     noplog.Logger,
		Output:  &target.Proto{},
	}

	file, err := os.Open("./fixtures/route_guide.proto")
	if err != nil {
		t.Fatalf("Failed to open proto file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			t.Errorf("error during closing file: %v", err)
		}
	}()

	protoFile := proto.NewParser(file)

	parsedProtoFile, err := protoFile.Parse()
	if err != nil {
		t.Fatalf("error parsing proto file: %v", err)
	}

	proto.Walk(
		parsedProtoFile,
		proto.WithPackage(d.PackageHandler),
		proto.WithService(d.ServiceHandler),
		proto.WithMessage(d.MessageHandler),
		proto.WithEnum(d.EnumHandler),
	)

	if d.Output.Name != "routeguide" {
		t.Errorf("expected name to be routeguide, got %s", d.Output.Name)
	}

	if len(d.Output.Services) != 1 {
		t.Errorf("expected 1 service, got %d", len(d.Output.Services))
	}

	if d.Output.Services[0].Name != "RouteGuide" {
		t.Errorf("expected service name to be RouteGuide, got %s", d.Output.Services[0].Name)
	}

	if len(d.Output.Services[0].RPCs) != 4 {
		t.Errorf("expected 4 RPCs, got %d", len(d.Output.Services[0].RPCs))
	}

	if len(d.Output.Messages) != 5 {
		t.Errorf("expected 5 messages, got %d", len(d.Output.Messages))
	}
}

func TestFixtures_Rce(t *testing.T) {
	var d = &handlers.Dependency{
		Verbose: true,
		Log:     noplog.Logger,
		Output:  &target.Proto{},
	}

	file, err := os.Open("./fixtures/rce.proto")
	if err != nil {
		t.Fatalf("Failed to open proto file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			t.Errorf("error during closing file: %v", err)
		}
	}()

	protoFile := proto.NewParser(file)

	parsedProtoFile, err := protoFile.Parse()
	if err != nil {
		t.Fatalf("error parsing proto file: %v", err)
	}

	proto.Walk(
		parsedProtoFile,
		proto.WithPackage(d.PackageHandler),
		proto.WithService(d.ServiceHandler),
		proto.WithMessage(d.MessageHandler),
		proto.WithEnum(d.EnumHandler),
	)

	if d.Output.Name != "rce" {
		t.Errorf("expected name to be 'rce', got %s", d.Output.Name)
	}

	if len(d.Output.Services) != 1 {
		t.Errorf("expected 1 service, got %d", len(d.Output.Services))
	}

	if d.Output.Services[0].Name != "CodeExecutionEngineService" {
		t.Errorf("expected service name to be 'CodeExecutionEngineService', got %s", d.Output.Services[0].Name)
	}

	if len(d.Output.Services[0].RPCs) != 3 {
		t.Errorf("expected 3 RPCs, got %d", len(d.Output.Services[0].RPCs))
	}

	if len(d.Output.Messages) != 7 {
		t.Errorf("expected 7 messages, got %d", len(d.Output.Messages))
	}

	// Each RPC must not have an empty name and fields
	// for their request and response
	for i, rpc := range d.Output.Services[0].RPCs {
		// We check request first
		if rpc.Request.Name == "" {
			t.Errorf("[#%d] expected RPC request name to be non-empty, got %s", i, rpc.Request.Name)
		}

		if rpc.Request.Name != "EmptyRequest" && len(rpc.Request.Fields) == 0 {
			t.Errorf("[#%d] expected RPC request fields to be non-empty for %s, got %d", i, rpc.Request.Name, len(rpc.Request.Fields))
		}

		if rpc.Response.Name == "" {
			t.Errorf("[#%d] expected RPC response name to be non-empty, got %s", i, rpc.Response.Name)
		}

		if len(rpc.Response.Fields) == 0 {
			t.Errorf("[#%d] expected RPC response fields to be non-empty, got %d", i, len(rpc.Response.Fields))
		}
	}
}
