package handlers

import (
	"strings"

	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

func (d *Dependency) MessageHandler(msg *proto.Message) {
	if d.Verbose {
		d.Log.Printf("found message: %s", msg.Name)
	}
	if d.collection == nil {
		d.collection = &collection{}
	}

	var comment string
	if msg.Comment != nil {
		comment = strings.Join(msg.Comment.Lines, "\n")
	}

	var message = target.Message{
		Name:    msg.Name,
		Comment: comment,
		Fields:  []target.Field{},
	}

	var currentCollection = &collection{
		Fields: []target.Field{},
	}

	for _, el := range msg.Elements {
		el.Accept(currentCollection)
	}

	for _, col := range currentCollection.Fields {
		message.Fields = append(message.Fields, target.Field{
			Name:     col.Name,
			Comment:  col.Comment,
			Type:     col.Type,
			Sequence: col.Sequence,
			Repeated: col.Repeated,
			Optional: col.Optional,
			Required: col.Required,
		})
	}

	d.collection.Messages = append(d.collection.Messages, message)
	d.Output.Messages = append(d.Output.Messages, message)

	for i, srv := range d.Output.Services {
		for j, rpc := range srv.RPCs {
			var requestTypeName string
			var responseTypeName string

			for _, r := range d.collection.RPC {
				if rpc.Name == r.Name {
					requestTypeName = r.RequestType
					responseTypeName = r.ResponseType
					break
				}
			}

			if requestTypeName == "" || responseTypeName == "" {
				// No request and/or response exists
				continue
			}

			for _, m := range d.collection.Messages {
				if m.Name == requestTypeName && d.Output.Services[i].RPCs[j].Request.Name == "" {
					d.Output.Services[i].RPCs[j].Request = m

					if d.Verbose {
						d.Log.Printf("found message %s for rpc %s", m.Name, d.Output.Services[i].RPCs[j].Name)
					}
				}

				if m.Name == responseTypeName && d.Output.Services[i].RPCs[j].Response.Name == "" {
					d.Output.Services[i].RPCs[j].Response = m

					if d.Verbose {
						d.Log.Printf("found message %s for rpc %s", m.Name, d.Output.Services[i].RPCs[j].Name)
					}
					break
				}
			}
		}
	}
}
