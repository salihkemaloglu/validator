package service

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/salihkemaloglu/validator/pkg/model"
)

// SoapCall does service call
func SoapCall(req *http.Request) (*model.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	r := &model.Response{}
	err = xml.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
