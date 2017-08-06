package xmltime

import (
	"encoding/xml"
	"regexp"
	"time"
)

func isWithoutZ(v string) bool {
	match, _ := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$", v)
	return match
}

// XMLTime comment
type XMLTime struct {
	time.Time
}

// UnmarshalXML comment
func (t *XMLTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	if isWithoutZ(v) {
		v += "Z"
	}
	// RFC3339: 2006-01-02T15:04:05Z07:00
	parse, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*t = XMLTime{parse}
	return nil
}

// UnmarshalXMLAttr comment
func (t *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	v := attr.Value
	if isWithoutZ(v) {
		v += "Z"
	}
	parse, err := time.Parse(time.RFC3339, attr.Value)
	if err != nil {
		return err
	}
	*t = XMLTime{parse}
	return nil
}
