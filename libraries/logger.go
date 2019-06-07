package libraries

import (
	"log"
	"strings"
	"unicode"
)

//remove whitespace
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func LogPanic(err error) {
	log.Panic("failed to load error configuration cause, %s", err.Error())
}
