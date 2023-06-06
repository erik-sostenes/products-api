package wrongs

// CommandNotRegisteredError will return an error when an command not registered
type CommandNotRegisteredError string

// CommandNotRegisteredError implements the Error interface
func (e CommandNotRegisteredError) Error() string {
	return string(e)
}

// CommandAlreadyRegisteredError will return an error when an command not registered
type CommandAlreadyRegisteredError string

// CommandAlreadyRegisteredError implements the Error interface
func (e CommandAlreadyRegisteredError) Error() string {
	return string(e)
}

// StatusBadRequest will return an error when the client makes a mistakes
type StatusBadRequest string

// StatusBadRequest implements the Error interface
func (e StatusBadRequest) Error() string {
	return string(e)
}
