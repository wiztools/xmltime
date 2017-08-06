package xmltime

import (
	"encoding/xml"
	"time"
)

// XMLTime comment
type XMLTime struct {
	time.Time
}

// UnmarshalXML comment
func (t *XMLTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = time.RFC3339 // 2006-01-02T15:04:05Z07:00
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*t = XMLTime{parse}
	return nil
}

// UnmarshalXMLAttr comment
func (t *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	parse, err := time.Parse(time.RFC3339, attr.Value)
	if err != nil {
		return err
	}
	*t = XMLTime{parse}
	return nil
}
