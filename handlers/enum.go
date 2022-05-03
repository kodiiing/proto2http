package handlers

import (
	"proto2http/target"
	"strings"

	"github.com/emicklei/proto"
)

func (d *Dependency) EnumHandler(en *proto.Enum) {
	d.Log.Printf("enum handler got called")
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

	d.collection.Enums = append(d.collection.Enums, enumItem)

}
