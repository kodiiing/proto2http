package handlers

import (
	"strings"

	"github.com/kodiiing/proto2http/target"

	"github.com/emicklei/proto"
)

func (d *Dependency) ServiceHandler(srv *proto.Service) {
	if d.collection == nil {
		d.collection = &collection{}
	}

	var comment string
	if srv.Comment != nil {
		comment = strings.Join(srv.Comment.Lines, "\n")
	}

	service := target.Service{
		Name:    srv.Name,
		Comment: comment,
		RPCs:    []target.RPC{},
	}

	for _, el := range srv.Elements {
		el.Accept(d.collection)
	}

	for _, col := range d.collection.RPC {
		service.RPCs = append(service.RPCs, target.RPC{
			Name:    col.Name,
			Comment: col.Comment,
		})
	}

	d.Output.Services = append(d.Output.Services, service)
}
