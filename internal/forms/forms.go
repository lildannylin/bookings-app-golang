package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

//Form creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// Valid return true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New ini a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// MinLength return false if field less than require
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d charaters long", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(filed string) {
	if !govalidator.IsEmail(f.Get(filed)) {
		f.Errors.Add(filed, "Invaid email address")
	}
}
