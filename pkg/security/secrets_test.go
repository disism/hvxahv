package security

import (
	"testing"
)

func TestGenPassword(t *testing.T) {
	r := GenPassword("hvxahv")
	t.Log(r)
}
