package handlers

import (
	"strings"

	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

func (d *Dependency) EnumHandler(en *proto.Enum) {
	if d.Verbose {
		d.Log.Printf("found enum: %s", en.Name)
	}
	var comment string
	if en.Comment != nil {
		comment = strings.Join(en.Comment.Lines, "\n")
	}

	tempCollection := &collection{EnumValues: []target.EnumValue{}}

	for _, el := range en.Elements {
		el.Accept(tempCollection)
	}

	var enumItem = target.Enum{
		Name:    en.Name,
		Comment: comment,
		Values:  tempCollection.EnumValues,
	}

	d.Output.Enums = append(d.Output.Enums, enumItem)
}
