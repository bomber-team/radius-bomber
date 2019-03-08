package radiusGenerator

import (
	"gitlab.com/radius-tank/radiusGenerator/models"
)

//Main generate algorithm todo
func generate(limit int) []models.Attribute {
	//parse
	//request users_script, get data from it
	//generate attributes not users attributes
	//merge datas
	//repeat for limit times
	//transfer to radius-attacker
	return []models.Attribute{}
}

//execute users script and parse the result todo best to move to another structure and use it
//todo maybe to add this method to model.Attribute and we can just call merge in? but we have to get something as json
func mergeWithusers(path string) []models.Attribute {
	//add parse
	return []models.Attribute{}
}
