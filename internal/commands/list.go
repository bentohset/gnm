package commands

import (
	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/storage"
	"gopkg.in/urfave/cli.v1"
)

func List(c *cli.Context, logger *logger.AppLogger, storage *storage.YAMLStorage) error {
	storage.Print()
	return nil
}
