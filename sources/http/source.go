package http

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
)

type Source struct {
	url string
}

func NewSource(config map[string]interface{}) (Source, error) {
	return Source{
		url: config["path"].(string),
	}, nil
}

func (s Source) GetKeys() (map[string]string, error) {
	keys := map[string]string{}

	resp, err := http.Get(s.url)
	if err != nil {
		return keys, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lineTokens := strings.SplitN(scanner.Text(), ",", 2)

		if len(lineTokens) != 2 {
			logrus.Warnf("Line: %s not comma separated", lineTokens)
			continue
		}

		keys[lineTokens[0]] = lineTokens[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return keys, nil
}
