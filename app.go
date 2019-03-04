package main

import "gitlab.com/radius-tank/radiusGenerator/parseAttribute"

func main() {
	parser := parseAttribute.Parser{
		Path: "examples/scenario1.txt",
	}

	parser.ReadFromFile()


}
