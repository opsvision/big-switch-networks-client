package bsn

type AuthResponse struct {
	Success       bool   `json:"success"`
	Error         string `json:"error"`
	SessionCookie string `json:"session_cookie"`
}
