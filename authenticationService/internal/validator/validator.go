package validator

import (
	"regexp"
	"slices"
)

var (
	EmailRX    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	UsernameRX = regexp.MustCompile(`^[A-Za-z]\w*[A-Za-z-0-9]$`)
	//PasswordRX = regexp.MustCompile(`^\d+[A-Za-z]+|[A-Za-z]+\d\w*$`)
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, msg string) {
	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = msg
	}
}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

func In(v string, list ...string) bool {
	return slices.Contains(list, v)
}

func Matches(v string, rx *regexp.Regexp) bool {
	return rx.MatchString(v)
}

func Unique(vals []string) bool {
	uniqueVals := make(map[string]struct{})

	for _, value := range vals {
		uniqueVals[value] = struct{}{}
	}

	return len(vals) == len(uniqueVals)
}
