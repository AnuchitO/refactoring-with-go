package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// START OMIT
type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

type errWrite struct {
	io.Writer // HL
	err       error
}

func (e *errWrite) Write(buf []byte) (int, error) { // HL
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
} // HL

// END OMIT

// WRITE OMIT
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWrite{Writer: w} // HL
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}

// END WRITE OMIT

func main() {
	var buf bytes.Buffer
	st := Status{Code: 555, Reason: "account not found."}
	headers := []Header{
		{"Content-Type", "application/json"},
	}
	body := strings.NewReader("this is a body")

	err := WriteResponse(&buf, st, headers, body)
	fmt.Printf("buf: \n%s\n", buf.String())
	fmt.Printf("error:", err)
}
