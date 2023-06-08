package jwt

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/services"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxe8efqXwuRZ4V1FoCmUkRln6loUfp8XtqdRWzJciVivkJZld
/jHONSqfTw4SpY0qQkD5G+uNYb3MSEOkBH9UCwCHoSdl1sjYDaWbx/Je4d/NM6YG
E3rsWgwvIcujGCLz3BQ2hO+57NqEAHQI3LkWXI+Rie4K5HQAVvMCat4UMe4CL++Z
52PC5yh1sgAbLVnoPBoLI7cYCERPhuexWm+gnjez1GKrWkvBACB2t+XAC+hU1TTc
Ry6aEHUvJ5+dlBglfPWT6UQJw7PJRaSELD8YLih4C0saxpq+M+ZMoriv1/phY/kk
g1ZafT+lydgBlMVxLnESqgm20zJqplsIg2X7EQIDAQABAoIBABtkooXIlW4oK/N5
srptkP2jiki2l9DyVZgBaRnbeMcQP/zsItQBNJarFW0td2suBEEzGMbCbMiwKct+
gP6WWJ1FL4AgIbn+BditqMedRYBhJtcVDRY5FujHcuZsdl/qxnEY4wq22rZq74XY
iTly7CNXQz8hkKRZYYqnCxibL5RRK6Z5Zo9ohTNFOZIw897k4s9FZEOmEG8pWho9
TNE9tgHBuuYQzaJjLgWqIM97TatDNv7KA996tFNNTmtPzQl1B4MFO5YdpL+OjJ2L
IQAzs8vD3fbCEpYb7UTWHCU8d9bpWpONByO8ZOo49wUg/bZIHxZwb5uQtg63E5z3
+gdmiFkCgYEA+A1m2tAMfTHScNSKy+pir9xHEjR8JdUABWeYmbtxES6Iu8uTy8Vy
i40J6YfgRtyzUqEeo18MKfJWzk11RBPrSO8uNFFpQPDoL3FcknFxbBlhmngpcAUG
vqYNCrVLyUNXI1vDXnFZW+eg6E9CKf1pJ1R9BPc/qALRDj/lzFxDW+sCgYEAzEah
5NVSD83+GlUd+t+Xtyf9UqkNyKzPLUm4On3mXVLmx35J4tJpihm5SHFUo9gxY3dn
JxU4vbxqFd6P8xCvWKEQShqrrlCHZ5YUnbHdTxEd+atKf64IMIfwX87P8Jb5kY1i
6gHbVP3i5X/HRLqz6prMpnOc6J0Xg+nw8eHpcfMCgYAliPmgcM0DANAETNU36B7I
179VbOXAX8viBXwc/zUr0WvVZwfVVOpxXYU7dlkkv+7OuRzGwfI4QriJ/USaaZ03
6yGFvy/7KLkpvLCyZEIyhmCznC1BCzGrFbtxfF+cc/kym4cjumk4NAOwQ5YSfosz
7WABqVxTkyGJU3f1hZyXwwKBgFA/x0Xwj8ZptFN/8MEnqaBoc1pP03xsdw9hkKBZ
6W/sK4FfmYMkChYYuPM+onOjcPOUas+txJa1OC/TOVXRzjDRRWb3R065kBgfm4W/
5CM1pEL7Cc9S/SCjpsjcpE/t36lQk/U+OX4QJ1zlb9EOT7PwkEkrzg6L+Dr4YpGD
oIQFAoGBANkKowkWZwTQjj+Rcsr2Tl99BWQouaX92hMYwHSLlV+IxA+gVCz3d5o8
XX6jW7U9o45KANn9TVbBV+TLR2iJyZWjUPgNSVRoPnlPJfFYFbIYbS9mnZ+x9JDf
Tkqn9t2JRnRp7OXIpdXPxe+glrrYZA365rb4aZ048RTT3omcREG4
-----END RSA PRIVATE KEY-----`

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxe8efqXwuRZ4V1FoCmUk
Rln6loUfp8XtqdRWzJciVivkJZld/jHONSqfTw4SpY0qQkD5G+uNYb3MSEOkBH9U
CwCHoSdl1sjYDaWbx/Je4d/NM6YGE3rsWgwvIcujGCLz3BQ2hO+57NqEAHQI3LkW
XI+Rie4K5HQAVvMCat4UMe4CL++Z52PC5yh1sgAbLVnoPBoLI7cYCERPhuexWm+g
njez1GKrWkvBACB2t+XAC+hU1TTcRy6aEHUvJ5+dlBglfPWT6UQJw7PJRaSELD8Y
Lih4C0saxpq+M+ZMoriv1/phY/kkg1ZafT+lydgBlMVxLnESqgm20zJqplsIg2X7
EQIDAQAB
-----END PUBLIC KEY-----`

func TestJWT_GenerateToken(t *testing.T) {
	tsc := []struct {
		Token[Claims]
		Claims
	}{
		{
			Token: NewJWT([]byte(privateKey), []byte(publicKey)),
			Claims: NewClaims(services.AccountResponse{
				AccountId:       "322d3103-6984-4d0b-bbd3-6efe3dc981ef",
				AccountUserName: "Erik",
			}),
		},
		{
			Token: NewJWT([]byte(privateKey), []byte(publicKey)),
			Claims: NewClaims(services.AccountResponse{
				AccountId:       "94343721-6baa-4cd5-a0b4-6c5d0419c02d",
				AccountUserName: "JaredNV",
			}),
		},
	}

	for i, ts := range tsc {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			token := ts.Token

			tokenString, err := token.Generate(ts.Claims)
			if err != nil {
				t.Errorf("error was expected %v, but error it was obtained %v", "null", err)
			}

			t.Logf("Token: %s", tokenString)
		})
	}
}

func TestJWT_ValidatorToken(t *testing.T) {
	tsc := []struct {
		Token[Claims]
		Claims
	}{
		{
			Token: NewJWT([]byte(privateKey), []byte(publicKey)),
			Claims: NewClaims(services.AccountResponse{
				AccountId:       "322d3103-6984-4d0b-bbd3-6efe3dc981ef",
				AccountUserName: "Erik",
			}),
		},
		{
			Token: NewJWT([]byte(privateKey), []byte(publicKey)),
			Claims: NewClaims(services.AccountResponse{
				AccountId:       "94343721-6baa-4cd5-a0b4-6c5d0419c02d",
				AccountUserName: "JaredNV",
			}),
		},
	}

	for i, ts := range tsc {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			token := ts.Token

			tokenString, err := token.Generate(ts.Claims)
			if err != nil {
				t.Skip(fmt.Sprintf("failure to generate a token: %s", err))
			}

			err = token.Validate(tokenString)
			if err != nil {
				t.Errorf("error was expected %v, but error it was obtained %v", "null", err)
			}
		})
	}
}
