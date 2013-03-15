package xmlparser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkXmlParser(b *testing.B) {

	b.StopTimer()

	input, err := ioutil.ReadFile("./inquiry.xml")

	if err != nil {
		fmt.Println("Couldn't read file", err.Error())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ParseXml(input)
	}
}
