package model

import "encoding/xml"

type Request struct {
	CountryCode string
	VATNumber   string
}

type Response struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope" json:"-"`
	SoapBody *SOAPBodyResponse
}

type SOAPBodyResponse struct {
	XMLName      xml.Name `xml:"Body" json:"-"`
	Resp         *CheckVatResponse
	FaultDetails *Fault
}

type Fault struct {
	XMLName     xml.Name `xml:"Fault" json:"-"`
	Faultcode   string   `xml:"faultcode" json:"faultCode"`
	Faultstring string   `xml:"faultstring" json:"faultString"`
}

type CheckVatResponse struct {
	XMLName     xml.Name `xml:"checkVatResponse" json:"-"`
	CountryCode string   `xml:"countryCode" json:"countryCode"`
	VatNumber   string   `xml:"vatNumber" json:"vatNumber"`
	RequestDate string   `xml:"requestDate" json:"requestDate"`
	Valid       string   `xml:"valid" json:"valid"`
}
