package rz

import "log"

type Rz struct {
	verbose bool
}

func New(verbose bool) *Rz {
	return &Rz{
		verbose,
	}
}

func (rz *Rz) log(format string, v ...any) {
	if rz.verbose {
		log.Default().Printf(format, v...)
	}
}
