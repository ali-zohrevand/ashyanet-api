package OutputAPI

type Message struct {
	Warning string `json:"warning"`
	Error   string `json:"error"`
	Info    string `json:"info"`
}
type Error struct {
	ErrorMessageList map[string]string `errors`
}
