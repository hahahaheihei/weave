package service

import (
	"fmt"
	"strconv"
	"time"
	"weave/pkg/model"

	"github.com/golang-jwt/jwt/v4"
)

const (
	Issuer = "weave.io"
	Secret = "weaveserver"
)

type CustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type JWTService struct {
	signKey        []byte
	issuer         string
	expireDuration int64
}

func NewJWTService() *JWTService {
	return &JWTService{
		signKey:        []byte(Secret),
		issuer:         Issuer,
		expireDuration: int64(24 * time.Hour.Seconds()),
	}
}

func (s *JWTService) CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			Name: user.Name,
			ID:   user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + s.expireDuration,
				NotBefore: time.Now().Unix() - 1000,
				Id:        strconv.Itoa(user.ID),
				Issuer:    s.issuer,
			},
		},
	)

	return token.SignedString(s.signKey)
}

func (s *JWTService) ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return s.signKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invaild token")
	}

	return &model.User{
		ID:   claims.ID,
		Name: claims.Name,
	}, nil
}
