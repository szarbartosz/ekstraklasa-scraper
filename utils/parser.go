package utils

import (
	"log"
	"strconv"
	"strings"
)

func ParseToInt(toBeParsed string) int {
	var parsed int
	var err error

	parsed, err = strconv.Atoi(toBeParsed)

	if err != nil {
		log.Panic("Error while parsing to int: ", err)
	}

	return parsed
}

func SanitizeString(input string) string {
	words := strings.Fields(input)
	noWhitespace := strings.Join(words, "")
	lowercase := strings.ToLower(noWhitespace)

	return lowercase
}
