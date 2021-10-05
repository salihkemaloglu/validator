package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/salihkemaloglu/validator/pkg/model/validate"
	"github.com/salihkemaloglu/validator/pkg/service"
	"github.com/salihkemaloglu/validator/pkg/util/convert"
)

// ValidateVatID validates vat-id for germany
func ValidateVatID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	vatİD := strings.TrimPrefix(r.URL.Path, "/validate/vat-id/")
	if err := validate.VATID(vatİD); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req, err := service.GenerateSOAPRequest(vatİD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := service.SoapCall(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.SoapBody.FaultDetails != nil {
		b, err := convert.ToJSON(resp.SoapBody.FaultDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, b, http.StatusInternalServerError)
		return
	}

	b, err := convert.ToJSON(resp.SoapBody.Resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%s", b)
}
