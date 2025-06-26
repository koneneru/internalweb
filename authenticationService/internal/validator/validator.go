package validator

import (
	"regexp"
	"slices"
)

var (
	EmailRX    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	UsernameRX = regexp.MustCompile(`^(?=.*[A-Za-z0-9]$)[A-Za-z][A-Za-z\d]{0,19}$`)
	// Minimum eight characters, at least one letter and one number:
	PasswordRX1 = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`)
	// Minimum eight characters, at least one letter, one number and one special character:
	PasswordRX2 = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)
	// Minimum eight characters, at least one uppercase letter, one lowercase letter and one number:
	PasswordRX3 = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`)
	// Minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character:
	PasswordRX4 = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	// Minimum eight and maximum 10 characters, at least one uppercase letter, one lowercase letter, one number and one special character:
	PasswordRX5 = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,10}$`)
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
