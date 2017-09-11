/*
Package xmltime provides struct XMLTime. XMLTime wraps `time.Time` type. XMLTime
can be used to convert date-time in RFC3339 format in XML element values / attributes
when unmarshalling.
*/
package xmltime

import (
	"encoding/xml"
	"time"
)

var raiseParseErrOnEmpty = true

func mustAddZ(v string) bool {
	_, err := time.Parse("2006-01-02T15:04:05", v)
	if err != nil {
		return false
	}
	return true
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

// UnmarshalXML implements xml.Unmarshaler interface.
func (t *XMLTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	return t.unmarshall(v)
}

// UnmarshalXMLAttr implements xml.UnmarshalerAttr interface.
func (t *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	v := attr.Value
	return t.unmarshall(v)
}

// MarshalXML implements xml.Marshaler interface.
func (t XMLTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(t.Format(time.RFC3339), start)
	return nil
}

// MarshalXMLAttr implements xml.MarshalerAttr interface.
func (t XMLTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name, Value: t.Format(time.RFC3339)}
	return attr, nil
}

var beginningTime = time.Time{}

// IsBeginning returns true if date-time is set to: 0001-01-01 00:00:00 +0000 UTC.
func (t *XMLTime) IsBeginning() bool {
	if t.Equal(beginningTime) {
		return true
	}
	return false
}
