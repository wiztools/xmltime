# Golang xmltime

Reference: https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields

## Introduction

Unmarshall XML element values and attributes to `time.Time` values.

To install:

```
$ go get github.com/wiztools/xmltime
```

Define the time data in XML elements as `xmltime.XMLTime` type, and when the data is unmarshalled from type ISO8601/RFC3339 date string, it is converted to `time.Time` type.

Refer the `xmltime_test.go` source code for example.
