package OutputAPI

type Message struct {
	Warning string
	Error   string
	Info    string
}
type Error struct {
	ErrorMessageList map[string]string `errors`
}
