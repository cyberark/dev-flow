package util

import (
	"fmt"
	"log"

	"strings"
)

func ValidateStringParam(paramName string, paramValue string, validValues []string) {
	valueFound := false

	for i := 0; i < len(validValues); i++ {
		valueFound = (paramValue == validValues[i])
		if valueFound { break }
	}

	if !valueFound {
		err := fmt.Sprintf(
			"Invalid value '%s' for param %s. Must be one of [%s].",
			paramValue,
			paramName,
			strings.Join(validValues, ", "),
		)
		log.Fatalln(err)		
	}
}
