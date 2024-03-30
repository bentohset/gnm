package commands

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/model"
	"github.com/bentohset/gnm/internal/storage"
	"github.com/bentohset/gnm/internal/utils"
	"gopkg.in/urfave/cli.v1"
)

func Create(c *cli.Context, logger *logger.AppLogger, storage *storage.YAMLStorage) error {
	host, createNew, err := promptFlow(c)
	if err != nil {
		logger.Error("[CREATE] Format error")
		color.Red("Invalid format: %s", err)
		return fmt.Errorf("invalid input")
	}

	if createNew {
		// ssh keygen with file name
		err = utils.Execute("", "ssh-keygen", "-f", host.PrivateKeyPath)
		if err != nil {
			color.Red("An error occured generating your keys. Check if ssh-keygen is installed")
			logger.Error("[CREATE] Create: unable to execute ssh-keygen")
			return err
		}

		// prompt to do ssh-copy-id to server
		isInsertStr := utils.GetInput("Auto copy the public key to server? y/n (y): ")
		isInsert, err := utils.ParseBoolean(isInsertStr)
		if err != nil {
			color.Red("Invalid input, should be y/n")
			logger.Error("[CREATE] Create: invalid input for auto copy")
			return fmt.Errorf("invalid input, should be y/n")
		}
		if isInsert {
			// copy the public key to server
			err = utils.Execute("", "ssh-copy-id")
			if err != nil {
				color.Red("An error occured copying your id to the remote server. Check if ssh-copy-id is installed")
				logger.Error("[CREATE] Create: unable to execute ssh-copy-id")
				return err
			}
		}
	}

	// save the config to hostfile
	_, err = storage.Save(*host)
	if err != nil {
		color.Red("An error occured saving your host to the config file")
		logger.Error("[CREATE] Create: Unable to save host to config file")
		return err
	}

	color.Green("Successfully created host config '%v'", host.Alias)

	return nil
}

// prompting flow for user input to create a host
func promptFlow(c *cli.Context) (*model.Host, bool, error) {
	h := model.Host{}

	h.User = utils.GetInput("Enter user: ")
	h.HostName = utils.GetInput("Enter hostname: ")
	h.Alias = utils.GetInput("Enter a unique alias for this connection (empty to use hostname): ")
	if h.Alias == "" {
		h.Alias = h.HostName
	}
	// TODO: check alias exists
	h.Description = utils.GetInput("Enter a short description (empty): ")
	h.PrivateKeyPath = utils.GetInput("Enter the path to the private key (empty to generate new): ")
	if h.PrivateKeyPath == "" {
		// generate new key path,
		keyFile := utils.GetInput("Enter the filename of the new keypair: ")
		if keyFile == "" {
			fmt.Println("Invalid input. Please enter a filename")
			return &h, true, fmt.Errorf("please enter a valid input")
		}
		sshPath := c.GlobalString("ssh-path")
		h.PrivateKeyPath = sshPath + "/" + keyFile
		return &h, true, nil
	}

	return &h, false, nil
}
