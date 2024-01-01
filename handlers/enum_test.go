package handlers_test

import (
	"testing"

	"github.com/kodiiing/proto2http/handlers"
	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

func TestEnumHandler(t *testing.T) {
	var d = &handlers.Dependency{
		Verbose: true,
		Log:     noplog.Logger,
		Output:  &target.Proto{},
	}

	var en = &proto.Enum{
		Name: "SampleEnum",
		Comment: &proto.Comment{
			Lines: []string{"Enum comment"},
		},
		Elements: []proto.Visitee{
			&proto.EnumField{
				Name: "Value1",
				Comment: &proto.Comment{
					Lines: []string{"Enum value comment"},
				},
				Integer: 1,
			},
			&proto.EnumField{
				Name:    "Value2",
				Integer: 2,
			},
		},
	}

	d.EnumHandler(en)

	if d.Output.Enums[0].Name != "SampleEnum" {
		t.Errorf("Expected enum name to be 'SampleEnum', got '%s'", d.Output.Enums[0].Name)
	}

	if d.Output.Enums[0].Comment != "Enum comment" {
		t.Errorf("Expected enum comment to be 'Enum comment', got '%s'", d.Output.Enums[0].Comment)
	}

	if len(d.Output.Enums[0].Values) != 2 {
		t.Errorf("Expected 2 enum values, got %d", len(d.Output.Enums[0].Values))
	}

	if d.Output.Enums[0].Values[0].Key != "Value1" {
		t.Errorf("Expected enum value name to be 'Value1', got '%s'", d.Output.Enums[0].Values[0].Key)
	}

	if d.Output.Enums[0].Values[0].Integer != 1 {
		t.Errorf("Expected enum value integer to be 1, got %d", d.Output.Enums[0].Values[0].Integer)
	}

	if d.Output.Enums[0].Values[0].Comment != "Enum value comment" {
		t.Errorf("Expected enum value comment to be 'Enum value comment', got '%s'", d.Output.Enums[0].Values[0].Comment)
	}

	if d.Output.Enums[0].Values[1].Key != "Value2" {
		t.Errorf("Expected enum value name to be 'Value2', got '%s'", d.Output.Enums[0].Values[1].Key)
	}

	if d.Output.Enums[0].Values[1].Integer != 2 {
		t.Errorf("Expected enum value integer to be 2, got %d", d.Output.Enums[0].Values[1].Integer)
	}

	if d.Output.Enums[0].Values[1].Comment != "" {
		t.Errorf("Expected enum value comment to be empty, got '%s'", d.Output.Enums[0].Values[1].Comment)
	}
}
