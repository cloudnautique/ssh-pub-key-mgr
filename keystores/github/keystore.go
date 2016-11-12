package gh

import (
	"crypto/md5"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh"
)

type KeyClient struct {
	client *github.Client
}

func NewKeyClient() (KeyClient, error) {
	return KeyClient{
		client: github.NewClient(nil),
	}, nil
}

func (kc KeyClient) GetKeyForUser(user string, wantedFingerprint string) (string, error) {
	logrus.Debugf("Looking for keys:  %s", user)

	opts := &github.ListOptions{PerPage: 100}
	keys, res, err := kc.client.Users.ListKeys(user, opts)
	if err != nil {
		return "", nil
	}

	logrus.Debugf("RESP:\n%#v", res)

	for _, key := range keys {
		logrus.Debugf("KEY: %v\n", key)

		fp, err := fingerprintMD5(*key.Key)
		if err != nil {
			return "", err
		}

		logrus.Debugf("FP: %s", fp)

		if fp == wantedFingerprint {
			logrus.Infof("Found key: %s", wantedFingerprint)
			return *key.Key, nil
		}
	}

	return "", fmt.Errorf("Key %s not found for user: %s", wantedFingerprint, user)
}

// There is a PR in the golang crypto lib that
// will handle this. For now.. its a copy
func fingerprintMD5(key string) (string, error) {
	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(key))
	if err != nil {
		logrus.Errorf("Couldn't parse pubKey: %#v", pubKey)
		return "", err
	}
	md5sum := md5.Sum(pubKey.Marshal())

	return rfc4716hex(md5sum[:]), nil
}

func rfc4716hex(data []byte) string {
	var out string
	for i := 0; i < len(data); i++ {
		out = fmt.Sprintf("%s%0.2x", out, data[i])
		if i != len(data)-1 {
			out = out + ":"
		}
	}
	return out
}
