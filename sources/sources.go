package sources

import (
	"github.com/cloudnautique/ssh-pub-key-mgr/sources/files"
)

type Source interface {
	GetKeys() (map[string]string, error)
}

func NewSource(source string) (Source, error) {
	return files.NewSource()
}
