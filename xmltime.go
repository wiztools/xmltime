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

// AllowEmptyDateTime empty dates to: 0001-01-01 00:00:00 +0000 UTC
func AllowEmptyDateTime() {
	raiseParseErrOnEmpty = false
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

// UnmarshalXMLAttr comment
func (t *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	v := attr.Value
	if mustAddZ(v) {
		v += "Z"
	}
	parse, err := time.Parse(time.RFC3339, attr.Value)
	if err != nil {
		return err
	}
	*t = XMLTime{parse}
	return nil
}

// IsBeginning comment
func (t *XMLTime) IsBeginning() bool {
	exp := time.Time{}
	if t.Equal(exp) {
		return true
	}
	return false
}
