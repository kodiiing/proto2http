package handlers

import (
	"strings"

	"github.com/emicklei/proto"
)

func (d *Dependency) PackageHandler(pkg *proto.Package) {
	d.Output.Name = pkg.Name

	var comment string
	if pkg.Comment != nil {
		comment = strings.Join(pkg.Comment.Lines, "\n")
	}

	d.Output.Comment = comment
}
