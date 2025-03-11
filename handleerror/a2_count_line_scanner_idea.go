package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// CountLines ...
// START OMIT
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r) // HL
	lines := 0

	for sc.Scan() { // HL
		lines++
	}
	return lines, sc.Err()
}

// END OMIT

func main() {
	s := `please
count
me
five
line`

	f := strings.NewReader(s)

	n, err := CountLines(f)
	fmt.Printf("n: %d, err: %v", n, err)
}
