package main

import (
	"os"
	"sort"

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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "source,s",
			Usage: "location to `file://PATH` or http(s)://URL containing allowed users and fingerprints",
		},
		cli.StringFlag{
			Name:  "keystore,k",
			Usage: "keystore backend",
			Value: "github",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}

func mainAction(c *cli.Context) error {
	logrus.Info("Calling Main")

	config, err := initConfig(c)
	if err != nil {
		return err
	}

	ds, err := sources.NewSource(config)
	if err != nil {
		logrus.Error(err)
		return err
	}

	keyClient, err := keystores.NewBackend(c.String("keystore"))
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
