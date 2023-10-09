package lo

import (
	"fmt"
	"os"
)

var DebugEnable bool

func init() {
	_, DebugEnable = os.LookupEnv("DEBUG")
}

func Debug(f string, v ...any) {
	if len(f) > 0 && f[len(f)-1] != '\n' {
		f += "\n"
	}
	if DebugEnable {
		fmt.Fprintf(os.Stderr, f, v...)
	}
}
