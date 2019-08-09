package testutils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJsonFile(path string) []byte {
	file, err := os.Open(path)
	defer file.Close()
	
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(file)
	
	return bytes
}
