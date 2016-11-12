package gh

import (
	"testing"
)

var testKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDIKEoxc7owerOglMlEprcv5WCXazCkUmgl86wQPwTyClYcxXTwMHnTc2rBm3kXpPxkbHp6RgQWoBqQ3PYRLSHCaz8UOQgrnLzHe7KHOubfM+32GiCvVd5s0lVQ5SVgR62CNauXnniDTVEPCHfOdlx2DB9Qh1Kpu8w+ahqJs88uU+CIHyHKMdMSGIyHGsv+pxX5DnbPNo9rFbUAW7dyHFReBCUK3BlApB29dvv0kwPkUWvK2M6p/2nRMkJ5rht0oi1is6OsmSU8ajun6wYh+Kd0qSJYlrCsNUu/ZS6WsBsRAnJA8dpmw36skrksvWMwQbJa22X+CgUTxp6h8bKqMh+h"

var fingerprint = "95:3e:5d:ce:7c:f8:40:e5:78:fc:ec:6f:68:53:4a:7e"

func TestMD5Fingerprinting(t *testing.T) {
	md5, err := fingerprintMD5(testKey)
	if md5 != fingerprint || err != nil {
		t.Errorf("MD5 doesnt match, got: %s should be %s", md5, fingerprint)
		t.Errorf("ERROR: %s", err)
	}
}
