package entity

type TokenError struct {
}

func (e TokenError) Error() string {
	return "token is not valid"
}
