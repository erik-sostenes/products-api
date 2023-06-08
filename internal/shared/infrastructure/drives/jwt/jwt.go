package jwt

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/services"
	"github.com/golang-jwt/jwt/v4"
)

type (
	// Claims interface which is a component of the jwt.Claims interface
	// Claims is the contract that every object must implement to generate the token payload
	Claims interface {
		jwt.Claims
	}

	// Token interface containing the methods to generate and validate a token,
	// the value of type any must be an object that implements the interface Claims
	Token[V any] interface {
		// Generate method that generates a token
		Generate(v V) (string, error)
		// Validate method that validates a token
		Validate(token string) error
	}
)

// claims is a structure representing the payload of a token
//
// claims composes the jwt.RegisteredClaims structure, which in turn implements the jwt.Claims interface
// in this way claims fulfills the contract to generate the token payload
type claims struct {
	ID    string
	User  string
	Email string
	jwt.RegisteredClaims
}

// NewClaims Instances a Claims with the data of an account
func NewClaims(account services.AccountResponse) Claims {
	return &claims{
		ID:   account.AccountId,
		User: account.AccountUserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 15)),
			Issuer:    "turing.ia",
		},
	}
}

// JWT representing all the business logic of a token using the standard jwt
// implements the Token[Claims] interface
type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewJWT returns an instance of the Token[Claims] interface
//
// NewJWT loads the public and private keys
func NewJWT(privKey, publKey []byte) Token[Claims] {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publKey)
	if err != nil {
		panic(err)
	}

	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// Generate method that generates and token with the claims to the payload and the private key
func (j *JWT) Generate(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

// Validate method that validates the token using the public key and the type of encryption method
func (j *JWT) Validate(token string) (err error) {
	_, err = jwt.Parse(token, j.validate)
	return
}

// validate method that represents a jwt.KeyFunc and its job is to validate the encryption type
// of the method if no error occurs the public key is returned
func (j *JWT) validate(token *jwt.Token) (any, error) {
	if method, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf(`invalid signing method %s`, method)
	}
	return j.publicKey, nil
}
