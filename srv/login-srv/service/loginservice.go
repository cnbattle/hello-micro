package loginservice

// LoginInterface interface
type LoginInterface interface {
	Login() (LoginRequest, error)
}

// LoginRequest LoginRequest
type LoginRequest struct {
	Token string
	//UID    string
	//IsAuth bool
}
