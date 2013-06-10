package utils

import (
  "testing"
)

func TestSha1(t *testing.T) {
  str := "bacaberita"
  sha1 := "2ea8729c122dfac9b7f2f283b80f1734603732f1"

  out := Sha1(str)
  if sha1 != out {
    t.Errorf("Invalid SHA1 for '%s' expected '%s' got '%s'",
             str, sha1, out)
  }
}

