package Helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/speps/go-hashids/v2"
)

var Salt string = "S4LT_C0D3_TH1NGK1"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func HashIdEncode(code int) string {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)

	id, _ := h.Encode([]int{code})

	return id
}

func HashIdDecode(code string) int {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeWithError(code)
	return d[0]
}
