package xmltime

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Dt      XMLTime  `xml:"dt"`
}

type RootEmpty struct {
	XMLName xml.Name `xml:"root"`
	Dt      XMLTime  `xml:"dt,omitempty"`
}

type RootAttr struct {
	XMLName xml.Name `xml:"root"`
	Dt      XMLTime  `xml:"dt,attr,omitempty"`
}

func getDt(t *testing.T, xmlData string) (XMLTime, error) {
	var env Root
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return env.Dt, err
	}
	return env.Dt, nil
}

func runTest(t *testing.T, xmlData string, exp string) {
	dt, err := getDt(t, xmlData)
	if err != nil {
		t.Fatal(err)
	}
	o := fmt.Sprintf("Dt: %v.", dt)
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

func TestEmptyXMLTime(t *testing.T) {
	emptyDtXML := "<root><dt/></root>"
	withDtXML := "<root><dt>2006-01-02T15:04:05Z</dt></root>"

	// Non-lenient test:
	dt, err := getDt(t, emptyDtXML)
	if err == nil {
		t.Fail()
	}

	// Lenient test with support for empty dates:
	AllowEmptyDateTime()
	runTest(t,
		emptyDtXML,
		"Dt: 0001-01-01 00:00:00 +0000 UTC.")
	dt, _ = getDt(t, emptyDtXML)

	// IsBeginning() test:
	if !dt.IsBeginning() {
		t.Fail()
	}

	dt, _ = getDt(t, withDtXML)
	if dt.IsBeginning() {
		t.Fail()
	}
}

func TestOmitEmptyXMLTime(t *testing.T) {
	emptyDtXML := "<root></root>"
	var env RootEmpty
	err := xml.Unmarshal([]byte(emptyDtXML), &env)
	if err != nil {
		t.Error(err)
	}
	if !env.Dt.IsBeginning() {
		t.Fail()
	}
}

func TestAttr(t *testing.T) {
	dtXML := "<root dt=\"2006-01-02T15:04:05Z\"></root>"
	var env RootAttr
	err := xml.Unmarshal([]byte(dtXML), &env)
	if err != nil {
		t.Error(err)
	}
	if env.Dt.Year() != 2006 {
		t.Fail()
	}
}

func TestMarshal(t *testing.T) {
	var env Root
	xmlData := "<root><dt>2006-01-02T15:04:05Z</dt></root>"
	_ = xml.Unmarshal([]byte(xmlData), &env)
	out, _ := xml.Marshal(env)
	if string(out) != xmlData {
		t.Fail()
	}
}

func TestMarshalAttr(t *testing.T) {
	var env RootAttr
	xmlData := "<root dt=\"2006-01-02T15:04:05Z\"></root>"
	_ = xml.Unmarshal([]byte(xmlData), &env)
	out, _ := xml.Marshal(env)
	if string(out) != xmlData {
		t.Fail()
	}
}
