package commands

import (
	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/storage"
	cli "gopkg.in/urfave/cli.v1"
)

func InitCommands(logger *logger.AppLogger, storage *storage.YAMLStorage) []cli.Command {
	return []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a new SSH key.",
			Action: func(c *cli.Context) error {
				return Create(c, logger, storage)
			},
		},
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "List all the available SSH keys",
			Action: func(c *cli.Context) error {
				return List(c, logger, storage)
			},
		},
		{
			Name:    "use",
			Aliases: []string{"u"},
			Usage:   "Set specific SSH key as default by its alias name",
			Action: func(c *cli.Context) error {
				return Use(c, logger, storage)
			},
		},
		{
			Name:    "del",
			Aliases: []string{"d"},
			Usage:   "Delete specific SSH key by alias name",
			Action: func(c *cli.Context) error {
				return Delete(c, logger, storage)
			},
		},
	}
}
