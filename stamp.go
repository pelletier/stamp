package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var parsedDateFormats = []string{
	"2006-01-02",
	"20060102",
	time.ANSIC,
}

func parseDate(input string) string {
	for _, format := range parsedDateFormats {
		parsedTime, err := time.Parse(format, input)
		if err == nil {
			return strconv.FormatInt(parsedTime.Unix(), 10)
		}
	}
	return ""
}

func parseUnixTimestamp(input string) string {
	integer, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return ""
	}
	parsedTime := time.Unix(integer, 0)
	return parsedTime.Format(parsedDateFormats[0])
}

func init() {
}

func main() {
	flag.Parse()
	inputs := flag.Args()
	for _, input := range inputs {
		result := parseDate(input)
		if len(result) == 0 {
			result = parseUnixTimestamp(input)
		}
		if len(result) == 0 {
			fmt.Println("Could not figure out what", input, "is")
		} else {
			fmt.Println(result)
		}
	}
}
