package types

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// type CustomClaims struct {
// 	Name     string `json:"name"`
// 	Picture  string `json:"picture"`
// 	Iss      string `json:"iss"`
// 	Aud      string `json:"aud"`
// 	AuthTime int64  `json:"auth_time"`
// 	UserId   string `json:"user_id"`
// 	Sub      string `json:"sub"`
// 	Iat      int64  `json:"iat"`
// 	Exp      int64  `json:"exp"`
// 	Email    string `json:"email"`
// 	jwt.StandardClaims
// }

type CustomClaims struct {
	LocalID          string `json:"localId"`
	Email            string `json:"email"`
	PhotoURL         string `json:"photoUrl"`
	EmailVerified    bool   `json:"emailVerified"`
	ProviderUserInfo []struct {
		ProviderID  string `json:"providerId"`
		PhotoURL    string `json:"photoUrl"`
		FederatedID string `json:"federatedId"`
		Email       string `json:"email"`
		RawID       string `json:"rawId"`
		ScreenName  string `json:"screenName"`
	} `json:"providerUserInfo"`
	ValidSince    string    `json:"validSince"`
	LastLoginAt   string    `json:"lastLoginAt"`
	CreatedAt     string    `json:"createdAt"`
	ScreenName    string    `json:"screenName"`
	LastRefreshAt time.Time `json:"lastRefreshAt"`
	jwt.StandardClaims
}
