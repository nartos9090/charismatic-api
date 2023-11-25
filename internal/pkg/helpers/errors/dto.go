package helpers_errors

type Error struct {
	HttpCode int
	Message  string
	Errors   []string
}
