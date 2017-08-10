/*
Package xmltime provides struct XMLTime. XMLTime wraps `time.Time` type. XMLTime
can be used to convert date-time in RFC3339 format in XML element values / attributes
when unmarshalling.
*/
package xmltime

import (
	"encoding/xml"
	"regexp"
	"time"
)

var raiseParseErrOnEmpty = true

func mustAddZ(v string) bool {
	if len([]rune(v)) == 19 {
		match, _ := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$", v)
		return match
	}
	return false
}

/*
AllowEmptyDateTime sets up the unmarshalling engine to convert empty dates in XML
to: 0001-01-01 00:00:00 +0000 UTC. In the absence of this flag being set, the parser
will terminate parsing and return error. This is a irreversible call, and a global one.
*/
func AllowEmptyDateTime() {
	raiseParseErrOnEmpty = false
}

// XMLTime wraps time.Time.
type XMLTime struct {
	time.Time
}

func (t *XMLTime) unmarshall(v string) error {
	if mustAddZ(v) {
		v += "Z"
	}
	if (!raiseParseErrOnEmpty) && v == "" {
		*t = XMLTime{time.Time{}}
		return nil
	}
	// RFC3339: 2006-01-02T15:04:05Z07:00
	parse, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*t = XMLTime{parse}
	return nil
}

// UnmarshalXML converts XML element value to xmltime.XMLTime.
func (t *XMLTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	return t.unmarshall(v)
}

// UnmarshalXMLAttr converts attribute value to xmltime.XMLTime.
func (t *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	v := attr.Value
	return t.unmarshall(v)
}

var beginningTime = time.Time{}

// IsBeginning returns true if date-time is set to: 0001-01-01 00:00:00 +0000 UTC.
func (t *XMLTime) IsBeginning() bool {
	if t.Equal(beginningTime) {
		return true
	}
	return false
}
