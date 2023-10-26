package main

import (
	"context"
	"flag"

	"github.com/aoissx/mcsrv/subcmd"
	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(subcommands.CommandsCommand(), "help")
	subcommands.Register(subcommands.FlagsCommand(), "help")
	subcommands.Register(subcommands.HelpCommand(), "help")

	// init command
	subcommands.Register(&subcmd.InitCmd{}, "")
	// start command
	subcommands.Register(&subcmd.StartCmd{}, "")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	subcommands.Execute(ctx)
}
