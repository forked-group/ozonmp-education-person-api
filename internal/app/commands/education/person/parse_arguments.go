package person

import (
	"errors"
	"strings"
	"unicode"
)

var ErrUnpairedQuotationMark = errors.New("unpaired quotation mark")

// NOTE: now "some text,"" other text" -> ["some text,", " other text"]

func parseArguments(text string) ([]string, error) {
	var args []string

	for len(text) != 0 {

		pos := len(text)
		for i, c := range text {
			if !unicode.IsSpace(c) {
				pos = i
				break
			}
		}

		if pos == len(text) {
			break
		}

		text = text[pos:]

		switch text[0] {
		case '"', '\'':
			q := text[0]
			text = text[1:]
			pos = strings.IndexByte(text, q)
			if pos == -1 {
				return nil, ErrUnpairedQuotationMark
			}
			args = append(args, text[:pos])
			pos++
		default:
			pos = len(text)
			for i, c := range text {
				if unicode.IsSpace(c) {
					pos = i
					break
				}
			}
			args = append(args, text[:pos])
		}

		text = text[pos:]
	}

	return args, nil
}
