package handlers_test

import (
	"testing"

	"github.com/kodiiing/proto2http/handlers"
	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

func TestPackageHandler(t *testing.T) {
	var d = &handlers.Dependency{
		Output: &target.Proto{},
	}

	var pkg = &proto.Package{
		Name: "SamplePackage",
		Comment: &proto.Comment{
			Lines: []string{"Package comment"},
		},
	}

	d.PackageHandler(pkg)

	if d.Output.Name != "SamplePackage" {
		t.Errorf("Expected package name to be 'SamplePackage', got '%s'", d.Output.Name)
	}

	if d.Output.Comment != "Package comment" {
		t.Errorf("Expected package comment to be 'Package comment', got '%s'", d.Output.Comment)
	}

}
