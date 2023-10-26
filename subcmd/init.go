package subcmd

import (
	"context"
	"flag"

	"github.com/aoissx/mcsrv/config"
	"github.com/google/subcommands"
)

type InitCmd struct{}

func (c *InitCmd) Name() string { return "init" }

func (c *InitCmd) Synopsis() string { return "Create config file." }

func (c *InitCmd) Usage() string { return "init" }

func (c *InitCmd) SetFlags(f *flag.FlagSet) {
}

func (c *InitCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	config.LogInfo("Create config file.")

	// Check config file exists.
	if config.CheckConfigFile() {
		config.LogError("Config file already exists.")
		config.LogError("If you want to reset the config file, please delete the config file.")
		return subcommands.ExitFailure
	}

	// Create default config file.
	err := config.SaveDefaultConfig()
	if err != nil {
		config.LogError("Failed to create config file.")
		return subcommands.ExitFailure
	}

	config.LogSuccess("Successfully created config file.")
	return subcommands.ExitSuccess
}
