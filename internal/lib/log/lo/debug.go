package lo

import (
	"github.com/rs/zerolog/log"
	"os"
)

var DebugEnable bool

func init() {
	_, DebugEnable = os.LookupEnv("DEBUG")
}

func Debug(f string, v ...any) {
	if DebugEnable {
		log.Printf(f, v...)
	}
}
