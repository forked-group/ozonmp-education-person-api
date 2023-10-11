package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var tmp map[string]any
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&tmp); err != nil {
		log.Fatalf("can't decode from stdin: %v", err)
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(tmp); err != nil {
		log.Fatalf("can't encode to stdout: %v", err)
	}
}
