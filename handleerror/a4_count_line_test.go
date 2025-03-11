package main

import (
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	f := strings.NewReader(`please
								count
								me
								five
								line`)
	n, err := CountLines(f)

	if n != 5 {
		t.Errorf("should count: 5 but got : %d\n", n)
	}

	if err != nil {
		t.Error("should not be an error but got :", err)
	}

}
