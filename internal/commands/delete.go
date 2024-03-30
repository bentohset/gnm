package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/storage"
	"github.com/bentohset/gnm/internal/utils"
	"gopkg.in/urfave/cli.v1"
)

func Delete(c *cli.Context, logger *logger.AppLogger, storage *storage.YAMLStorage) error {
	var alias string

	// input validation
	if c.NArg() > 0 {
		alias = c.Args().Get(0)
	} else {
		color.Red("Please provide an input alias name")
		logger.Error("[DELETE] Input key alias name not given")
		return fmt.Errorf("please input key alias name")
	}

	// check if alias exists
	host, err := storage.Get(alias)
	if err != nil {
		color.Red("Alias '%v' does not exist", alias)
		logger.Error("[DELETE] Alias '%v' does not exist", alias)
		return err
	}

	// prompt for user confirmation
	isDeleteStr := utils.GetInput("Confirm deletion of config and keys? y/n (n): ")
	isDelete, err := utils.ParseBoolean(isDeleteStr)
	if err != nil {
		color.Red("Please provide 'y' or 'n' only")
		logger.Error("[DELETE] Delete: invalid input for confirm deletion")
		return err
	}
	if !isDelete {
		return nil
	}

	// delete key files
	os.Remove(host.PrivateKeyPath)
	os.Remove(host.PrivateKeyPath + ".pub")

	// delete key from host config
	err = storage.Delete(host.Alias)
	if err != nil {
		color.Red("An error occured deleting your host config")
		logger.Error("[DELETE] Delete: unable to delete key from host config")
		return err
	}

	color.Green("Successfully deleted alias '%v'", alias)

	return nil
}
