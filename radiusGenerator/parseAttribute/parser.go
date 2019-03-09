package parseAttribute

import (
	"bufio"
	"fmt"
	"gitlab.com/radius-tank/radiusGenerator/models"
	"log"
	"os"
)

type Parser struct {
	Path string
}

//Check that file available
func (parser *Parser) check(e error) {
	if e != nil {
		panic(e)
	}
}

//read file by lines
func (parser *Parser) ReadFromFile() (attributes []models.ScenarioAttribute) {
	scenario, err := os.Open(parser.Path)
	parser.check(err)

	defer scenario.Close()

	scanner := bufio.NewScanner(scenario)
	automat := CheckAutomat{}

	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
		attributes = append(attributes, automat.checkFormatAndGet(scanner.Bytes()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return attributes
}
