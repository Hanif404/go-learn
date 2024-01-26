package generate

import (
	"testing"
)

func TestGenerateNIP(t *testing.T) {
    want := 0
	msg, err := GenerateNIP("ikhwan", 2019, 2, 1)
    if len(msg) == want || err != nil {
        t.Fatalf(`data not match, error : %v`, err)
    }
}

func TestGenerateNextNIP(t *testing.T) {
    want := 0
	msg, err := GenerateNextNIP("ART191-00002", 10)
    if len(msg) == want || err != nil {
        t.Fatalf(`data not match, error : %v`, err)
    }
}