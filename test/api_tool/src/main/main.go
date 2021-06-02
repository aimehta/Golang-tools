package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Purpose: Read the json file and generate unit tests
//

func getit(jsonFileName string) (bool, int) {
	fpath, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(fpath)
	}
	fmt.Println(fpath)

	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result)
	return true, 0
}

func main() {

	getit("refrence-01.json")
}
