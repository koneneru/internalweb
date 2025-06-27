package data

import (
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func Win1251ToUTF8(s string) (string, error) {
	dec := charmap.Windows1251.NewDecoder()
	buf, err := io.ReadAll(dec.Reader(strings.NewReader(s)))
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
