package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadToJson(path string) (result map[string]interface{}, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//var anyJson map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	println(result)
	defer jsonFile.Close()
	return result, nil
}
