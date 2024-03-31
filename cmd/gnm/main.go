package main

import (
	"log"
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/bentohset/gnm/internal/commands"
	"github.com/bentohset/gnm/internal/constants"
	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/storage"
)

var home, _ = os.UserHomeDir()
var defaultSSHPath = filepath.Join(home, ".ssh")

func main() {
	var err error
	// get app dir
	appDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("[MAIN] Can't get application home folder: %v", err)
	}
	// init logger
	lg, err := logger.New(appDir, "info")
	if err != nil {
		log.Fatalf("[MAIN] Can't get application home folder: %v", err)
	}
	// init storage
	strg, err := storage.NewYAML(appDir, &lg)
	if err != nil {
		lg.Error("[MAIN] Error running application: %v", err)
		os.Exit(1)
	}
	// check if ssh installed

	parseArgs()

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ssh-path",
			Value: defaultSSHPath,
			Usage: "Path to a .ssh folder",
		},
	}

	app.Name = constants.Name
	app.Usage = constants.Usage
	app.Version = constants.Version
	app.Commands = commands.InitCommands(&lg, strg)

	if err := app.Run(os.Args); err != nil {
		lg.Error("[MAIN] Can't save application state before closing %v", err)
	}
}

func parseArgs() {
	if len(os.Args) == 1 {
		constants.PrintLogo()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "h" || os.Args[1] == "help" {
			constants.PrintLogo()
		}
	}
}
