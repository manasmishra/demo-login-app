package common

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "app.rsa.pub"
)

// private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the keys before starting http handler
func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

// Generate JWT Token
func GenerateJWT(name, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// Middleware for validating JWT Tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// SetCors(&w)
	// Get token from request
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	}, request.WithClaims(&AppClaims{}))
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(
					w,
					err,
					"Access Token is expired get a new Token",
					401,
				)
				return
			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the access token",
					500,
				)
				return
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing Access Token!",
				500,
			)
			return
		}
	}
	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
