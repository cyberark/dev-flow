package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Confirm(prompt string) bool {
	r := bufio.NewReader(os.Stdin)

	tries := 3
	
	for ; tries > 0; tries-- {
		fmt.Printf("%v [y/n]: ", prompt)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}
