package handlers

import (
	"log"

	"github.com/kodiiing/proto2http/target"
)

type Dependency struct {
	Verbose bool
	Log     *log.Logger
	Output  *target.Proto

	collection *collection
}
