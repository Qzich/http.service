package validation

import (
	"regexp"
	"strings"

	"github.com/Qzich/http.service"
)

type StringField struct {
	http_service.JSONString
}

//returns sanitized value
func (value StringField) Value() string {
	return strings.TrimSpace(value.JSONString.Value)
}

func (value StringField) IsNotEmpty() bool {
	return value.Value() != ""
}

func (value StringField) IsMatchFormat(r *regexp.Regexp) bool {
	return r.MatchString(value.Value())
}
