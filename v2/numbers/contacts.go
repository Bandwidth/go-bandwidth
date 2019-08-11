package numbers

import "encoding/xml"

// Contact for an Account.
type Contact struct {
	XMLName   xml.Name `xml:"Contact"`
	FirstName string   `xml:"FirstName"`
	LastName  string   `xml:"LastName"`
	Phone     string   `xml:"Phone"`
	Email     string   `xml:"Email"`
}
