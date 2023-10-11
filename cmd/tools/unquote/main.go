package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	buf, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("can't read from stdin: %v", err)
	}

	lastQuote := bytes.LastIndexByte(buf, '"')
	n := 0
	for i := 0; i <= lastQuote; i++ {
		if buf[i] == '\n' || buf[i] == '\r' {
			continue
		}
		buf[n] = buf[i]
		n++
	}
	buf = buf[:n]

	s, err := strconv.Unquote(string(buf))
	if err != nil {
		log.Fatalf("can't unquote text: %v", err)
	}

	_, err = w.WriteString(s)
	if err != nil {
		log.Fatalf("can't write to stdout: %v", err)
	}
}
