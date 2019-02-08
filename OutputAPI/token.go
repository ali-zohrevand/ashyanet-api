package OutputAPI

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}
type TokenValid struct {
	Valid bool `json:"valid"`
}
