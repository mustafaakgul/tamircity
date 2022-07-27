package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)
//Read json file and return map slice
func ReadToJson(path string) (result *map[string]interface{}, error)
{
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	defer jsonFile.Close()
	return &result, nil
}