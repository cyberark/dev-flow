package util

import (
	"fmt"
	"log"
)

func ValidateStringParam(paramName string, paramValue string, validValues []string) {
	validValueMap := make(map[string]bool)
	for _, validValue := range validValues {
		validValueMap[validValue] = true
	}
	
	if !validValueMap[paramValue] {
		err := fmt.Sprintf(
			"Invalid value '%s' for param %s. Must be one of %v.",
			paramValue,
			paramName,
			validValues,
		)
		log.Fatalln(err)		
	}
}
