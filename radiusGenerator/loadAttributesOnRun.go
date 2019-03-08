package radiusGenerator

import (
	"layeh.com/radius"
	"layeh.com/radius/dictionary"
	"log"
	"strconv"
)

var Dictionary = map[string]radius.Type{}

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

//load 'attributes from transfer file to use them on run
func (loader *Loader) ParseFile() {
	parser := &dictionary.Parser{
		Opener: &dictionary.FileSystemOpener{},
	}

	for _, path := range loader.Paths {
		dict, err := parser.ParseFile(path)
		if err != nil {
			log.Fatal(err)
		}
		loader.addToDictionaryStandartAttributes(dict.Attributes)
		//Add vendors to our memory
		for _, element := range dict.Vendors {
			loader.addToDictionaryVendorJradius(element)
		}
	}
}

//load standard attributes to our dictionary
func (loader Loader) addToDictionaryStandartAttributes(attributes []*dictionary.Attribute) {
	for _, attrubute := range attributes {
		i, err := strconv.Atoi(attrubute.OID.String())
		if err != nil {
			panic(err)
		}
		Dictionary[attrubute.Name] = radius.Type(i)
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
		Dictionary[element.Name] = radius.Type(dict.Number<<16 + i)
	}
}
