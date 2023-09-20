package lo

import (
	"log"
	"os"
)

var DebugEnable bool

func init() {
	if _, ok := os.LookupEnv("DEBUG"); ok {
		DebugEnable = true
	}
}

func Debug(f string, v ...any) {
	if DebugEnable {
		if n := len(f); n == 0 || f[n-1] != '\n' {
			// Да я знаю, что логер умеет сам добавлять конец строки.
			// Да здесь пробел пред концом строки не лишний.
			// В Goland в выводе тестов между стоками лога появлялись пустые стоки.
			// Подбешивало. Если добавить конец строки таким образом, то вроде не появляются...
			// ???: ХЗ на что это влияет, но тфу-тфу-тфу работает.

			f += " \n"
		}
		log.Printf(f, v...)
	}
}
