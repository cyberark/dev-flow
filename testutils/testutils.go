package testutils

import (
	"encoding/json"
	
	"io/ioutil"
	"os"
)

func loadJsonFile(path string) []byte {
	file, err := os.Open(path)
	defer file.Close()
	
	if err != nil {
		panic("Failed to load fixture.")
	}

	bytes, _ := ioutil.ReadAll(file)
	
	return bytes
}

func LoadFixture(path string, obj interface{}) {
	bytes := loadJsonFile(path)

	err := json.Unmarshal(bytes, obj)

	if err != nil {
		panic("Failed to unmarshal json.")
	}
}
