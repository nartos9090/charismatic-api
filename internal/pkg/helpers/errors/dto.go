package errors

type Error struct {
	HttpCode int
	Message  string
	Errors   []string
}
