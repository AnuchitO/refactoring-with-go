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
	var (
		br    = bufio.NewReader(r) // HL
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF { // HL
		return 0, err
	}

	return lines, nil
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
