package forms

import (
	"net/http"
	"net/url"
)

// Form creaters a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// Valid checks if there's no error
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}


// New creates a form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has check if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}
	return true
}
