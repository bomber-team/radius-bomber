package main

import (
	"encoding/json"
	"gitlab.com/radius-tank/radiusAttacker"
	"gitlab.com/radius-tank/radiusGenerator"
	"layeh.com/radius"
	"os"
	"strconv"
)

type Application struct {
	Address           string
	Limit             string
	Secret            string
	PathsDictionaries []string
	ScenarioPath      string
}

var app = Application{}

//First we need to load dictionary in memory
//Next we start our cycle dos
func main() {
	errSettings := loadSettings("")
	if errSettings != nil {
		panic(errSettings)
	}

	err := loadDictionary()
	if err != nil {
		panic(err)
	}

	generator := radiusGenerator.Generator{
		app.ScenarioPath,
	}
	number, errConvert := strconv.Atoi(app.Limit)
	if errConvert != nil {
		panic(errConvert)
	}
	packets := generator.Generate(number)

	attacker := radiusAttacker.Attacker{
		radius.CodeAccessRequest,
		[]byte(app.Secret),
		app.Address,
	}
	attacker.Attack(packets)
}

func loadSettings(path string) error {
	//filename is the path to the json config file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&app)
	if err != nil {
		return err
	}
	return err
}

func loadDictionary() error {
	loader := radiusGenerator.Loader{}
	return loader.ParseFile()
}
