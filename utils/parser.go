package utils

import (
	"log"
	"strconv"
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
