package subcmd

import (
	"context"
	"flag"

	"github.com/aoissx/mcsrv/config"
	"github.com/aoissx/mcsrv/download"
	"github.com/google/subcommands"
)

type StartCmd struct {
	skip bool
}

func (c *StartCmd) Name() string { return "start" }

func (c *StartCmd) Synopsis() string { return "Start Server." }

func (c *StartCmd) Usage() string { return "start" }

func (c *StartCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.skip, "skip", false, "Skip Download Server Jar.")
}

func (c *StartCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	config.LogInfo("Start Server.")

	conf, err := config.GetConfig()
	if err != nil {
		config.LogError("Failed to get config.")
		return subcommands.ExitFailure
	} else {
		config.LogSuccess("Successfully got config.")
	}

	if !c.skip {
		// Download Server Jar.
		err := download.DownloadServerJar(conf)
		if err != nil {
			config.LogError("Failed to download server jar.")
			return subcommands.ExitFailure
		}
		config.LogSuccess("Successfully downloaded server jar.")
	}

	//todo: start server
	return subcommands.ExitSuccess
}
