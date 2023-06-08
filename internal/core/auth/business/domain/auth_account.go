package domain

// AuthAccount (Value Object) represents the auth account
type AuthAccount struct {
	authId       AuthId
	authPassword AuthPassword
}

// NewAuthAccount returns an instance of AuthAccount, receives and valid the primitive
// values of auth id and auth password
func NewAuthAccount(id string, password string) (authAccount AuthAccount, err error) {
	authId, err := NewAuthID(id)
	if err != nil {
		return
	}

	authPassword, err := NewAuthPassword(password)
	if err != nil {
		return
	}
	return AuthAccount{
		authId:       authId,
		authPassword: authPassword,
	}, nil
}

// PasswordMatches check the match of passwords
func (a AuthAccount) PasswordMatches(password AuthPassword) error {
	return a.authPassword.Equals(password)
}

func (a AuthAccount) AuthId() AuthId {
	return a.authId
}
