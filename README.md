# Golang xmltime

Unmarshall XML element values and attributes in ISO8601/RFC3339 date-time format to `xmltime.XMLTime` values. `xmltime.XMLTime` just wraps `time.Time`:

```go
type XMLTime struct {
    time.Time
}
```

The library is inspired from [this SO question and the answers suggested therein](https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields).

To install:

```
$ go get github.com/bambooengg/xmltime
```

or, if you use [dep](https://github.com/golang/dep):

```
$ dep ensure github.com/bambooengg/xmltime
```

## Parsing

Take this example XML:

```xml
<root>
  <dt>2006-01-02T15:04:05</dt>
</root>
```

To unmarshall this XML, a struct like this needs to be defined:

```go
import "encoding/xml"
import "github.com/bambooengg/xmltime"

type Root struct {
	XMLName xml.Name `xml:"root"`
	Dt      xmltime.XMLTime  `xml:"dt"`
}
```

Unmarshalling, and getting a `xmltime.XMLTime` representation of the data is as simple as:

```go
import "fmt"
import "encoding/xml"

var d Root
// load XML into byte[] b
err := xml.Unmarshal(b, &d)
// handle error
fmt.Printf("Date: %v.\n", d.Dt)
```

Refer the `xmltime_test.go` source code for example.

## Empty date-time values

When empty date-time values are encountered in XML, there are two ways to handle it:

1. Unmarshall parse error: This will terminate the XML parsing.
2. Assign a beginning date (since it's NOT possible to assign `nil` to struct instances in Go).

The default behavior of this library is to error out the parsing, which is the option 1 above. If you want to use the library as described in option 2, you need to set an internal flag to the package:

```go
xmltime.AllowEmptyDateTime()
```

This is an irreversible call, and a global one. So use it with caution. Once this is set, empty dates will be converted to `0001-01-01 00:00:00 +0000 UTC`, hence forth called the **beginning date**. We also provide a convenience function to validate if a date value is converted to beginning date. Continuing from our example above:

```go
if !d.Dt.isBeginning() {
    // do your stuff!
}
```
