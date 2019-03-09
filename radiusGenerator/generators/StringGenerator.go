package generators

type StringGenerator struct{}

//generate with concrete length string
func (StringGenerator) GenerateString(length int) string {
	return ""
}

//Generate default length random string
func (StringGenerator) GenerateStringDefault() string {
	return string("")
}
