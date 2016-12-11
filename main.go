package main

import (
	"fmt"
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
	app.Usage = "set source flag, and send in the path to write authorized_keys"
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
	authorizedKeys := []string{}

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

	for user, fingerprints := range keys {
		keySet, err := keyClient.GetKeysForUser(user, fingerprints)
		if err != nil {
			return err
		}
		authorizedKeys = append(authorizedKeys, keySet...)
	}

	return writeAuthorizedKeysFile(true, authorizedKeys, config["authorizedKeysPath"].(string))
}

func writeAuthorizedKeysFile(clobber bool, keys []string, file string) error {
	tmpFile := fmt.Sprintf("%s/authorized_keys.tmp", file)

	f, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = os.Chmod(tmpFile, 0600)
	if err != nil {
		return err
	}

	for _, key := range keys {
		f.WriteString(fmt.Sprintf("%s\n", key))
	}

	err = os.Rename(tmpFile, fmt.Sprintf("%s/authorized_keys", file))
	if err != nil {
		return err
	}

	return nil
}
