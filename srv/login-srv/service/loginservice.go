package loginservice

type LoginInterface interface {
	Login() (LoginRequest, error)
}

type LoginRequest struct {
	Token string
	//Uid    string
	//IsAuth bool
}
