package handlers

import (
	"log"
	"proto2http/target"
)

type Dependency struct {
	Verbose bool
	Log     *log.Logger
	Output  *target.Proto

	collection *collection
}
