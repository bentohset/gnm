package commands

import (
	"fmt"

	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/storage"
	"github.com/bentohset/gnm/internal/utils"
	"github.com/fatih/color"
	"gopkg.in/urfave/cli.v1"
)

func Use(c *cli.Context, logger *logger.AppLogger, storage *storage.YAMLStorage) error {
	var alias string
	if c.NArg() > 0 {
		alias = c.Args().Get(0)
	} else {
		color.Red("Please provide an input alias name")
		logger.Error("[USE] Use: Input key alias name not given")
		return fmt.Errorf("please input key alias name")
	}

	// check if alias exists
	host, err := storage.Get(alias)
	if err != nil {
		color.Red("Alias '%v' does not exist", alias)
		logger.Error("[USE] Use: Alias '%v' does not exist", alias)
		return err
	}

	// execute ssh
	remote := host.User + "@" + host.HostName
	err = utils.Execute("", "ssh", remote, "-i", host.PrivateKeyPath)
	if err != nil {
		color.Red("An error occured trying to ssh into remote server")
		logger.Error("[USE] Use: unable to execute ssh")
		return err
	}

	return nil
}
