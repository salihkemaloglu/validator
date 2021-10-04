package service

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"net/http"

	"github.com/salihkemaloglu/validator/pkg/model"
)

func getTemplate() string {
	return `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
	xmlns:tns1="urn:ec.europa.eu:taxud:vies:services:checkVat:types"
	xmlns:impl="urn:ec.europa.eu:taxud:vies:services:checkVat">
	<soap:Header>
	</soap:Header>
	<soap:Body>
		<tns1:checkVat xmlns:tns1="urn:ec.europa.eu:taxud:vies:services:checkVat:types"
		xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
		<tns1:countryCode>{{.CountryCode}}</tns1:countryCode>
		<tns1:vatNumber>{{.VATNumber}}</tns1:vatNumber>
		</tns1:checkVat>
	</soap:Body>
	</soap:Envelope>`
}

// GenerateSOAPRequest generates soap request
func GenerateSOAPRequest(vatID string) (*http.Request, error) {
	template, err := template.New("InputRequest").Parse(getTemplate())
	if err != nil {
		return nil, err
	}

	req := model.Request{
		CountryCode: vatID[:2],
		VATNumber:   vatID[2:],
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, "http://ec.europa.eu/taxation_customs/vies/services/checkVatService", bytes.NewBuffer(doc.Bytes()))
	if err != nil {
		return nil, err
	}

	return r, nil
}
