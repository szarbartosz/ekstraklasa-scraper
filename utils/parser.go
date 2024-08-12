package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var months = map[string]time.Month{
	"stycznia":     time.January,
	"lutego":       time.February,
	"marca":        time.March,
	"kwietnia":     time.April,
	"maja":         time.May,
	"czerwca":      time.June,
	"lipca":        time.July,
	"sierpnia":     time.August,
	"września":     time.September,
	"października": time.October,
	"listopada":    time.November,
	"grudnia":      time.December,
}

func ParseToInt(toBeParsed string) int {
	var parsed int
	var err error

	sanitized := SanitizeString(toBeParsed)

	parsed, err = strconv.Atoi(sanitized)

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

func ParseDateTime(timeString, dateString string) (time.Time, error) {
	if dateString == "" {
		return time.Time{}, nil
	}

	dateWords := strings.Fields(dateString)

	if len(dateWords) < 4 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateString)
	}

	day := ParseToInt(dateWords[1])
	month := months[dateWords[2]]
	year := ParseToInt(dateWords[3])

	var hour, minute int
	if timeString != "" {
		if _, err := fmt.Sscanf(timeString, "%d:%d", &hour, &minute); err != nil {
			return time.Time{}, err
		}
	}

	return time.Date(year, month, day, hour, minute, 0, 0, time.UTC), nil
}
