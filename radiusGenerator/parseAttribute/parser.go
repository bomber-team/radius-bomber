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

//read file by lines todo
func (parser *Parser) ReadFromFile() []models.ScenarioAttribute {
	var attributes []models.ScenarioAttribute
	scenario, err := os.Open(parser.Path)
	parser.check(err)

	defer scenario.Close()

	scanner := bufio.NewScanner(scenario)

	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
		attributes = append(attributes, parser.checkFormat(scanner.Bytes()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return attributes
}

const (
	SPACE              = byte(' ')
	GenStateKey        = byte('G')
	ValueStateKey      = byte('V')
	IntegratedStateKey = byte('I')
)

const (
	StartState      = 1
	TagState        = 2
	GenState        = 3
	ValueState      = 4
	IntegratedState = 5
)

//Automate for format checking todo
func (parser *Parser) checkFormat(s []byte) models.ScenarioAttribute {
	var currentState = StartState
	for _, c := range s {
		switch currentState {
		case StartState:
			switch c {
			case SPACE:
				fmt.Print("Space")
				currentState = TagState
			}
		case TagState:
			switch c {
			case GenStateKey:
				fmt.Print("GEN")
				currentState = GenState
			case ValueStateKey:
				fmt.Println("VALUE")
				currentState = ValueState
			case IntegratedStateKey:
				fmt.Println("Integrated")
				currentState = IntegratedState
			}
		case GenState:
			switch c {

			}
		case ValueState:
			switch c {

			}
		case IntegratedState:
			switch c {

			}
		}
		switch c {
		case SPACE:
			fmt.Print("Space")
		case GenStateKey:
			fmt.Print("Gen")
		case ValueStateKey:
			fmt.Print("VALUE")
		case IntegratedStateKey:
			fmt.Print("INTEGRATED")
		}
	}
	return models.ScenarioAttribute{}
}

func makeAttrbutePacket() {
	//packet := radius.New(radius.CodeAccessRequest, []byte(`secret`))

	//packet.Add(radius.Type(1), "")
}
