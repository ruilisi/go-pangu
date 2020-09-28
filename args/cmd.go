package args

import (
	"flag"
)

// CmdArgs ...
type CmdArgs struct {
	DB      string
	GIN_ENV string
	TABLE   string
}

var Cmd CmdArgs

func ParseCmd() {
	flag.StringVar(&Cmd.DB, "db", "", "db operation")
	flag.StringVar(&Cmd.GIN_ENV, "env", "development", "gin environment")
	flag.StringVar(&Cmd.TABLE, "table", "", "table name")
	flag.Parse()
}
