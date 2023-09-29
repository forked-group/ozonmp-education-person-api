package person

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode"
)

var ErrUnclosedQuotationMark = errors.New("unclosed quotation mark detected")

func splitIntoArguments(s string) ([]string, error) {
	var args []string

	r := strings.NewReader(s)
	var buf bytes.Buffer

	for r.Len() != 0 {
		c, _, err := r.ReadRune()

		for err == nil && unicode.IsSpace(c) {
			c, _, err = r.ReadRune()
		}

		if err == io.EOF {
			break
		}

		for err == nil && !unicode.IsSpace(c) {
			switch c {
			default:
				buf.WriteRune(c)
			case '"', '\'':
				q := c
				c, _, err = r.ReadRune()

				for err == nil && c != q {
					buf.WriteRune(c)
					c, _, err = r.ReadRune()
				}

				if err == io.EOF {
					args = append(args, string(buf.Bytes()))
					return args, ErrUnclosedQuotationMark
				}
			}

			c, _, err = r.ReadRune()
		}

		if err != nil && err != io.EOF {
			return nil, err
		}

		args = append(args, string(buf.Bytes()))
		buf.Reset()
	}

	return args, nil
}
