package lightning

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	if _, err := NewCollection("1234"); err != nil {
		t.Error("Collections are broke.")
	}
}
