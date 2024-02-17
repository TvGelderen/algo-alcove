package types

import "regexp"

type RegisterParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterErrors struct {
	Email    string
	Password string
	Other    string
}

func (params RegisterParams) Validate() (RegisterErrors, bool) {
	errors := RegisterErrors{}
	hasErrors := false

	if params.Email == "" {
		errors.Email = "Email is required"
		hasErrors = true
	} else {
		emailValid, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, params.Email)
		if !emailValid {
			errors.Email = "Invalid email address"
			hasErrors = true
		}
	}

	if params.Password == "" {
		errors.Password = "Password is required"
		hasErrors = true
	}

	return errors, hasErrors
}
