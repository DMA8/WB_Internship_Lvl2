package pkg

import (
	"log"
	"regexp"
)

func InitRegexCompile(cmd CmdLine) *regexp.Regexp {
	var regEx string
	if cmd.Fixed {
		regEx = cmd.RawPattern
	} else {
		regEx = cmd.Pattern
	}
	pattern, err := regexp.Compile(regEx)
	if err != nil {
		log.Fatal(err)
	}
	return pattern
}
