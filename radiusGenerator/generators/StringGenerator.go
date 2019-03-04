package generators

type StringGenerator struct{}

//generate with concrete length string
func (StringGenerator) generateString(length int) string {
	return ""
}

//Generate default length random string
func (StringGenerator) generateStringDefault() string {
	return string("")
}
