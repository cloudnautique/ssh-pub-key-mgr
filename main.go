package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/cloudnautique/ssh-pub-key-mgr/keystores"
	"github.com/cloudnautique/ssh-pub-key-mgr/sources"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "ssh-pub-key-mgr"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = mainAction

	app.Run(os.Args)
}

func mainAction(c *cli.Context) error {
	logrus.Info("Calling Main")

	ds, err := sources.NewSource("file")
	if err != nil {
		return err
	}

	keyClient, err := keystores.NewBackend("github")
	if err != nil {
		return err
	}
	keys, err := ds.GetKeys()
	if err != nil {
		return err
	}

	for user, fingerprint := range keys {
		keyClient.GetKeyForUser(user, fingerprint)
	}

	return nil
}
