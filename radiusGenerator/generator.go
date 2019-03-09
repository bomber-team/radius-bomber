package radiusGenerator

import (
	"context"
	"gitlab.com/radius-tank/radiusGenerator/generators"
	"gitlab.com/radius-tank/radiusGenerator/models"
	"gitlab.com/radius-tank/radiusGenerator/parseAttribute"
)

type Generator struct {
	Path string
}

//Main generate algorithm
func (generator Generator) generate(limit int) (packets [][]models.Attribute) {
	//parse
	parser := parseAttribute.Parser{
		generator.Path,
	}

	scenarioAttributes := parser.ReadFromFile()

	for i := 0; i < limit; i++ {
		packets = append(packets, generator.generateAttributes(scenarioAttributes))

	}

	return
}

///Generate attributes for each tag, //todo for users attributes there are two ways the first we get users script,
//todo merge with ours attributes and change state and next continue to fill packets
//todo second way made a structure which run users script once and contains attributes, when we need such attributes
//todo we ask structure to give them
func (generator Generator) generateAttributes(scenarioAttributes []models.ScenarioAttribute) (result []models.Attribute) {
	macGenerator := generators.MacGenerator{}
	stringGenerator := generators.StringGenerator{}

	for _, sAttribute := range scenarioAttributes {
		current := models.Attribute{
			Name: sAttribute.Name,
		}
		switch sAttribute.Tag {
		case parseAttribute.IntegratedStateKey:
			context.TODO()
		case parseAttribute.GenStateKey:
			switch sAttribute.Value {
			case "MAC":
				current.Value = macGenerator.GenRandomMac()
			case "NUMBER":

			case "WORD":
				current.Value = stringGenerator.GenerateStringDefault()
			}
		case parseAttribute.ValueStateKey:
			current.Value = sAttribute.Value
		}
		result = append(result, current)
	}
	return
}
