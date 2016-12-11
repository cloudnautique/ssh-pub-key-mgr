package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
)

type Source struct {
	filePath string
}

func NewSource(config map[string]interface{}) (Source, error) {
	path := strings.SplitN(config["path"].(string), "://", 2)

	return Source{
		filePath: path[1],
	}, nil
}

func (s Source) GetKeys() (map[string]string, error) {
	keys := map[string]string{}

	file, err := os.Open(s.filePath)
	if err != nil {
		return keys, err
	}

	scanner := bufio.NewScanner(file)
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
