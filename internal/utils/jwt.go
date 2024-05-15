package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceI interface {
	GeneraleToken(userid uint64) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTAuthCustomClaims struct {
	Userid uint64 `json:"userid"`
	// User     bool   `json:"user"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secretKey string   // 签名密钥
	issure    string   // 签发者
	subject   string   // 签发对象
	id        string   // wt ID, 类似于盐值
	audience  []string // 签发受众
}

func NewJwtGeneraleService() *JwtService {
	return &JwtService{
		secretKey: getSecretKey(),
		issure:    "haoyunpan-api2",
		subject:   "app",
		id:        RandStr(32, 5),
		audience:  jwt.ClaimStrings{"Android", "IOS"},
	}
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("secret")
	if secret == "" {
		secret = "secret"
	}

	return secret
}

// 登录时生成
func (service *JwtService) GeneraleToken(userid uint64) string {
	claims := &JWTAuthCustomClaims{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 最早使用时间
			Issuer:    service.issure,                                     // 签发者
			Subject:   service.subject,                                    // 签发对象
			ID:        service.id,                                         // wt ID, 类似于盐值
			Audience:  service.audience,                                   // 签发受众
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	singedStr, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}

	return singedStr
}

func (service *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})

	switch {
	case token.Valid:
		return token, nil
		// fmt.Println("You look nice today")
	case errors.Is(err, jwt.ErrTokenMalformed):
		return nil, errors.New("that's not even a token")
		// fmt.Println("That's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, errors.New("invalid signature")
		// Invalid signature
		// fmt.Println("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, errors.New("timing is everything")
		// Token is either expired or not active yet
		// fmt.Println("Timing is everything")
	default:
		// fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}
}

func (service *JwtService) ValidateWithClaimsToken(encodedToken string) (*jwt.Token, error) {
	claims := &JWTAuthCustomClaims{}
	token, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})

	// 获取userid
	// userid := claims.Userid

	switch {
	case token.Valid:
		return token, nil
		// fmt.Println("You look nice today")
	case errors.Is(err, jwt.ErrTokenMalformed):
		return nil, errors.New("that's not even a token")
		// fmt.Println("That's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, errors.New("invalid signature")
		// Invalid signature
		// fmt.Println("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, errors.New("timing is everything")
		// Token is either expired or not active yet
		// fmt.Println("Timing is everything")
	default:
		// fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}
}
