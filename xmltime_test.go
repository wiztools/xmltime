package xmltime

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Dt      XMLTime  `xml:"dt"`
}

func runTest(t *testing.T, xmlData string, exp string) {
	var env Root
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		log.Fatal("Error unmarshilling: ", err)
	}
	o := fmt.Sprintf("Dt: %v.", env.Dt)
	// fmt.Println(o)
	if o != exp {
		t.Fail()
	}
}

func TestXMLTime(t *testing.T) {
	runTest(t,
		"<root><dt>2006-01-02T15:04:05</dt></root>",
		"Dt: 2006-01-02 15:04:05 +0000 UTC.")
	runTest(t,
		"<root><dt>2006-01-02T15:04:05Z</dt></root>",
		"Dt: 2006-01-02 15:04:05 +0000 UTC.")
	runTest(t,
		"<root><dt>2006-01-02T15:04:05+05:30</dt></root>",
		"Dt: 2006-01-02 15:04:05 +0530 IST.")
}
