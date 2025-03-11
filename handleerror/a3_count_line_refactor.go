package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Scanner struct {
	br  *bufio.Reader
	err error
}

func (s *Scanner) Scan() bool {
	if s.err != nil {
		return false
	}

	s.read()

	return true
}

func (s *Scanner) Error() error {
	if s.err == io.EOF {
		return nil
	}

	return s.err
}

func (s *Scanner) read() {
	_, s.err = s.br.ReadString('\n')
}

func New(r io.Reader) *Scanner {
	return &Scanner{br: bufio.NewReader(r)}
}

func CountLines(r io.Reader) (int, error) {
	var lines int
	sc := New(r)
	for sc.Scan() {
		lines++
	}

	return lines, sc.Error()
}

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
