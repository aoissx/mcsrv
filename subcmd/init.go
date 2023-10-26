package subcmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type InitCmd struct{}

func (c *InitCmd) Name() string { return "init" }

func (c *InitCmd) Synopsis() string { return "Create config file." }

func (c *InitCmd) Usage() string { return "init" }

func (c *InitCmd) SetFlags(f *flag.FlagSet) {
}

func (c *InitCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
