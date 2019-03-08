package main

import "gitlab.com/radius-tank/radiusGenerator/parseAttribute"

//First we need to load dictionary in memory
//Next we start our cycle dos
func main() {
	parser := parseAttribute.Parser{
		Path: "examples/scenario1.txt",
	}

	parser.ReadFromFile()

}
