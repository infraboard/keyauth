package auth

type AuthCodeRequest struct {
	Code  string
	State string
}
