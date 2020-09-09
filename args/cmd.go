package args

import (
	"flag"
)

// CmdArgs ...
type CmdArgs struct {
	DB      string
	Amqp    int
	GIN_ENV string
}

var Cmd CmdArgs

func ParseCmd() {
	flag.StringVar(&Cmd.DB, "db", "", "db operation")
	flag.Parse()
}
