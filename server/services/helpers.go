package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"srv/structs"

	"github.com/go-playground/validator/v10"
)

//ParseJsonConfig - parsing json file after fing and validating file
func ParseJsonConfig() (structs.Config, []string) {
	cfg := structs.Config{}
	validate := validator.New()
	file, _ := os.Open("config.json")
	bytesJson, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytesJson, &cfg)
	errorsFields := []string{}

	if errors := validate.Struct(cfg); errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			errorsFields = append(errorsFields, err.Field())
		}
		return structs.Config{}, errorsFields
	}

	return cfg, nil
}
