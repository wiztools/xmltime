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

func TestXMLTime(t *testing.T) {
	xmlData := "<root><dt>2006-01-02T15:04:05Z</dt></root>"
	var env Root
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		log.Fatal("Error unmarshilling: ", err)
	}
	o := fmt.Sprintf("Dt: %v.", env.Dt)
	fmt.Println(o)
	if o != "Dt: 2006-01-02 15:04:05 +0000 UTC." {
		t.Fail()
	}
}
