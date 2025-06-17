package utils

import (
	"errors"
	"testing"
)

func TestShouldPanicOnErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Must("val", errors.New("test error"))
}
