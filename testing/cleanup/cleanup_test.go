package cleanup

import (
	"os"
	"testing"
)

func setup(filename string) func() {
	teardown := func() {
		err := os.Remove(filename)
		if err != nil {
			// panic("could not delete file")
		}
	}

	// create file ...

	return teardown
}

func TestSomething(t *testing.T) {
	teardown := setup("filename")
	defer teardown()

	t.Skip("TODO: implement.")
}
