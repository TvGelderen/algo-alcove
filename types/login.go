package types

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginErrors struct {
	Email    string
	Password string
	Other    string
}

func (params LoginParams) Validate() (LoginErrors, bool) {
	errors := LoginErrors{}
	hasErrors := false

	if params.Email == "" {
		errors.Email = "Email is missing"
		hasErrors = true
	}
	if params.Password == "" {
		errors.Password = "Password is missing"
		hasErrors = true
	}

	return errors, hasErrors
}
