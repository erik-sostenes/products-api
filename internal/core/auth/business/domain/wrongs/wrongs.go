package wrongs

type InvalidAuthAccount string

// InvalidAuthAccount implements the Error interface
func (e InvalidAuthAccount) Error() string {
	return string(e)
}

// InvalidAuthCredentials valid that the credentials of the account are correct
type InvalidAuthCredentials string

// InvalidAuthCredentials implements the Error interface
func (e InvalidAuthCredentials) Error() string {
	return string(e)
}
