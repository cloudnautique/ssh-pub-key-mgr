package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func initConfig(c *cli.Context) (map[string]interface{}, error) {
	config := map[string]interface{}{}

	if c.String("source") == "" {
		return config, errors.New("Need a valid source")
	}

	// Sources are in the form: <backend>://<path>
	source := strings.SplitN(c.String("source"), ":", 2)

	if len(source) == 2 {
		config["source"] = source[0]
		config["path"] = c.String("source")
	} else {
		return config, fmt.Errorf("Invalid config line: %s", c.String("source"))
	}

	config["refresh"] = c.Int("refresh-interval")

	if len(c.Args()) <= 0 {
		return config, fmt.Errorf("No path to write authorized_keys file given")
	}

	config["authorizedKeysPath"] = c.Args().Get(0)

	return config, nil
}
