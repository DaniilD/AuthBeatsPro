package JWT

import (
	"AuthBeatsPro/internal/configs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"strconv"
	"time"
)

type JWTHelper struct {
	jwtConfig *configs.JWTConfig
}

func NewJWTHelper(jwtConfig *configs.JWTConfig) *JWTHelper {
	return &JWTHelper{
		jwtConfig: jwtConfig,
	}
}

func (helper *JWTHelper) CreateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix(),
		Subject:   strconv.Itoa(userId),
	})

	tokenString, err := token.SignedString([]byte(helper.jwtConfig.SigningKey))

	return tokenString, err
}

func (helper *JWTHelper) NewRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(bytes)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}

func (helper *JWTHelper) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(helper.jwtConfig.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
