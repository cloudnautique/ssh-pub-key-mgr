package keystores

import (
	"github.com/cloudnautique/ssh-pub-key-mgr/keystores/github"
)

type BackendKeyClient interface {
	GetKeyForUser(string, string) (string, error)
}

func NewBackend(backendType string) (BackendKeyClient, error) {
	return gh.NewKeyClient()
}
