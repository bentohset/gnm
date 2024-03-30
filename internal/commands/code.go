package commands

import (
	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/storage"
	"gopkg.in/urfave/cli.v1"
)

// ref: https://code.visualstudio.com/docs/remote/ssh
func Code(c *cli.Context, logger *logger.AppLogger, storage *storage.YAMLStorage) error {
	// check if "code" can execute

	// get the alias

	// check if config is in .ssh/config file

	// else add it in
	// Host = alias
	// HostName = hostname
	// User = user
	// IdentityFile = idfilepath

	// run code --remote ssh-remote+<alias>

	return nil
}
