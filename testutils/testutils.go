package testutils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJsonFile(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(file)
	
	return string(bytes)
}
