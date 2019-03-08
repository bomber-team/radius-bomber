package tests

import (
	"gitlab.com/radius-tank/radiusGenerator"
	"testing"
)

func TestParse(t *testing.T) {
	slice := [1] string{"/home/kostya05983/go/src/gitlab.com/radius-tank/examples/testDictionary.txt"}

	loader := radiusGenerator.Loader{
		Paths: slice[0:1],
	}

	loader.ParseFile()
	l := radiusGenerator.Dictionary
	print(l)
}

func BenchmarkParse(t *testing.B) {

	//loader := radiusGenerator.Loader{
	//	"../examples/testDictionary.txt",
	//}
	//loader.ParseFile()
	//println()
}
