// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package words

import (
	"bufio"
	"bytes"
	"io"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

type Dict map[rune][]string

func (d Dict) Choice(word []byte) []byte {
	for i := len(word); i > 0; {
		r, size := utf8.DecodeLastRune(word[:i])

		if row, ok := d[unicode.ToUpper(r)]; ok {
			return []byte(row[rand.Intn(len(row))])
		}

		i -= size
	}

	return nil
}

func LoadDict(r io.Reader, dict Dict) {
	scanner := bufio.NewScanner(r)
	wordbuf := bytes.NewBuffer(nil)

	for scanner.Scan() {
		text := scanner.Text()

		if Normalize(wordbuf, text) {
			word := wordbuf.String()
			r, _ := utf8.DecodeRuneInString(word)

			dict[r] = append(dict[r], word)
		}
	}
}

func Normalize(buf *bytes.Buffer, word string) bool {
	buf.Reset()

	r, size := utf8.DecodeRuneInString(word)
	buf.WriteRune(unicode.ToUpper(r))

	for _, r := range word[size:] {
		if !unicode.IsLetter(r) {
			return false
		}

		buf.WriteRune(unicode.ToLower(r))
	}

	return true
}
