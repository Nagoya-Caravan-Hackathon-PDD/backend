package middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/firebase"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	AuthorizationHeaderKey = "Authorization"
	AuthorizationType      = "Bearer"
	PayloadContextKey      = "payload"
)

func (m *middleware) FirebaseAuth(n echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("FirebaseAuth")

		authHeader := c.Request().Header.Get(AuthorizationHeaderKey)
		if authHeader == "" {
			log.Println("Authorization header is not found")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 {
			log.Println("Authorization header is invalid")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		if authHeaderParts[0] != AuthorizationType {
			log.Println("Authorization header type is invalid")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		td := &TokenDecoder{tokenString: authHeaderParts[1]}
		header, err := td.Decode()
		if err != nil {
			log.Println("Failed to decode token")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		kid := header["kid"].(string)
		certString := firebase.GoogleJWks[kid].(string)

		cp := &CertificateParser{certString: certString}
		publicKey, err := cp.Parse()
		if err != nil {
			log.Println("Failed to parse certificate")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		tv := &TokenVerifier{tokenString: authHeaderParts[1], publicKey: publicKey}
		claims, err := tv.Verify()
		if err != nil {
			log.Println("Failed to verify token")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		c.Set(PayloadContextKey, claims)
		return n(c)
	}
}

type TokenDecoder struct {
	tokenString string
}

func (td *TokenDecoder) Decode() (map[string]interface{}, error) {
	parts := strings.Split(td.tokenString, ".")
	headerJson, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	var header map[string]interface{}
	err = json.Unmarshal(headerJson, &header)
	if err != nil {
		return nil, err
	}
	return header, nil
}

type CertificateParser struct {
	certString string
}

func (cp *CertificateParser) Parse() (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(cp.certString))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert.PublicKey.(*rsa.PublicKey), nil
}

type TokenVerifier struct {
	tokenString string
	publicKey   *rsa.PublicKey
}

func (tv *TokenVerifier) Verify() (*types.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tv.tokenString, &types.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tv.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*types.CustomClaims); ok && token.Valid {
		if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
			return nil, errors.New("Token is valid. But token is expired.")
		} else {
			return claims, nil
		}
	} else {
		return nil, errors.New("Token is not valid")
	}
}
