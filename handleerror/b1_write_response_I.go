package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

// START OMIT
func WriteResponse(w io.Writer, status Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", status.Code, status.Reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}

// END OMIT

// MAIN OMIT
func main() {
	var buf bytes.Buffer
	st := Status{Code: 555, Reason: "account not found."}
	headers := []Header{
		{"Content-Type", "application/json"},
	}
	body := strings.NewReader("this is a body")

	err := WriteResponse(&buf, st, headers, body)

	fmt.Printf("buf: \n%s\n", buf.String())
	fmt.Println("error:", err)
}

// ENDMAIN OMIT
