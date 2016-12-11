package sources

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/cloudnautique/ssh-pub-key-mgr/sources/files"
	"github.com/cloudnautique/ssh-pub-key-mgr/sources/http"
)

type Source interface {
	GetKeys() (map[string]string, error)
}

func NewSource(opts map[string]interface{}) (Source, error) {
	sourceType := opts["source"]

	switch sourceType {
	case "file":
		return files.NewSource(opts)
	case "http", "https":
		return http.NewSource(opts)
	default:
		logrus.Fatalf("Unknown source type: %s", sourceType)
	}

	return nil, fmt.Errorf("Unknonw type")
}
