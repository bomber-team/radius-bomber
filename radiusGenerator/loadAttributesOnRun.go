package radiusGenerator

import (
	"gitlab.com/radius-tank/radiusGenerator/models"
	"layeh.com/radius"
	"layeh.com/radius/dictionary"
	"log"
	"strconv"
)

//Global dictionary for easy parse user scenarious
var Dictionary = map[string]models.AttributeParams{}

//Structure for parse radius attribute on run
type Loader struct {
	Paths []string
}

//Checks that file opens correctly
func (loader Loader) check(e error) {
	if e != nil {
		panic(e)
	}
}

//load 'attributes from transfer file to use them on run todo made normal errors from function
func (loader *Loader) ParseFile() (err error) {
	parser := &dictionary.Parser{
		Opener: &dictionary.FileSystemOpener{},
	}
	for _, path := range loader.Paths {
		dict, err := parser.ParseFile(path)
		if err != nil {
			log.Fatal(err)
		}
		loader.addToDictionaryStandardAttributes(dict.Attributes)
		//Add vendors to our memory
		for _, element := range dict.Vendors {
			loader.addToDictionaryVendorJradius(element)
		}
	}
	return
}

//load standard attributes to our dictionary
func (loader Loader) addToDictionaryStandardAttributes(attributes []*dictionary.Attribute) {
	for _, attrubute := range attributes {
		i, err := strconv.Atoi(attrubute.OID.String())
		if err != nil {
			panic(err)
		}
		Dictionary[attrubute.Name] = models.AttributeParams{
			Type: attrubute.Type,
			OID:  radius.Type(i),
		}
	}
}

//parse our attribute shift them with jradius compatibility
//Add them to dictionary
func (loader *Loader) addToDictionaryVendorJradius(dict *dictionary.Vendor) {
	for _, element := range dict.Attributes {
		i, err := strconv.Atoi(element.OID.String())
		if err != nil {
			panic(err)
		}
		Dictionary[element.Name] = models.AttributeParams{
			Type: element.Type,
			OID:  radius.Type(dict.Number<<16 + i),
		}
	}
}
