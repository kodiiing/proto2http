package handlers

import (
	"strings"

	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

type temporaryRpc struct {
	Name         string
	Comment      string
	RequestType  string
	ResponseType string
}

type collection struct {
	proto.Visitor
	RPC        []temporaryRpc
	Messages   []target.Message
	Fields     []target.Field
	Enums      []target.Enum
	EnumValues []target.EnumValue
	Comments   []string
}

// //VisitProto(p *Proto)
// VisitMessage(m *Message)
// VisitService(v *Service)
// VisitSyntax(s *Syntax)
// VisitPackage(p *Package)
// VisitOption(o *Option)
// VisitImport(i *Import)
// VisitNormalField(i *NormalField)
// VisitEnumField(i *EnumField)
// VisitEnum(e *Enum)
// VisitComment(e *Comment)
// VisitOneof(o *Oneof)
// VisitOneofField(o *OneOfField)
// VisitReserved(r *Reserved)
// VisitRPC(r *RPC)
// VisitMapField(f *MapField)
// // proto2
// VisitGroup(g *Group)
// VisitExtensions(e *Extensions)

func (c *collection) VisitRPC(r *proto.RPC) {
	var comment string
	if r.Comment != nil {
		comment = strings.Join(r.Comment.Lines, "\n")
	}

	c.RPC = append(c.RPC, temporaryRpc{
		Name:         r.Name,
		Comment:      comment,
		RequestType:  r.RequestType,
		ResponseType: r.ReturnsType,
	})
}

func (c *collection) VisitMessage(m *proto.Message) {
	var comment string
	if m.Comment != nil {
		comment = strings.Join(m.Comment.Lines, "\n")
	}

	var message = target.Message{
		Name:    m.Name,
		Comment: comment,
		Fields:  []target.Field{},
	}

	var currentCollection = &collection{
		Fields: []target.Field{},
	}

	for _, el := range m.Elements {
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

	c.Messages = append(c.Messages, message)
}

func (c *collection) VisitNormalField(e *proto.NormalField) {
	var comment string
	if e.Comment != nil {
		comment = strings.Join(e.Comment.Lines, "\n")
	}

	c.Fields = append(c.Fields, target.Field{
		Name:     e.Name,
		Comment:  comment,
		Type:     e.Type,
		Sequence: int16(e.Sequence),
		Repeated: e.Repeated,
		Required: e.Required,
		Optional: e.Optional,
	})
}

func (c *collection) VisitMapField(f *proto.MapField) {
	var comment string
	if f.Comment != nil {
		comment = strings.Join(f.Comment.Lines, "\n")
	}

	c.Fields = append(c.Fields, target.Field{
		Name:     f.Name,
		Comment:  comment,
		Type:     f.Type,
		Sequence: int16(f.Sequence),
		Repeated: false,
	})
}

func (c *collection) VisitEnum(e *proto.Enum) {
	var comment string
	if e.Comment != nil {
		comment = strings.Join(e.Comment.Lines, "\n")
	}

	c.Enums = append(c.Enums, target.Enum{
		Name:    e.Name,
		Comment: comment,
	})
}

func (c *collection) VisitEnumField(i *proto.EnumField) {
	var comment string
	if i.Comment != nil {
		comment = strings.Join(i.Comment.Lines, "\n")
	}

	c.EnumValues = append(c.EnumValues, target.EnumValue{
		Key:     i.Name,
		Comment: comment,
		Integer: int16(i.Integer),
	})
}
