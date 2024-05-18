package middleware

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

func NewAuthInterceptor() grpc.UnaryServerInterceptor {
	return auth.UnaryServerInterceptor(authFunc)
}

type UserClaims struct {
	Address    string `json:"address,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	jwt.StandardClaims
}

var jwtSecret = os.Getenv("JWT_SECRET")

// var jwtSecret = "JWT_SECRET"

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(jwtSecret))
}

func ParseAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}

type JWTAuth struct {
}

// AuthFuncOverride is called instead of exampleAuthFunc.
func (s *JWTAuth) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	log.Info().Msgf("client is calling method: %s", fullMethodName)
	return ctx, nil
}

var tokenInfoKey struct{}

// authFunc is used by a middleware to authenticate requests
func authFunc(ctx context.Context) (context.Context, error) {
	if next, err := Next(ctx); err != nil {
		return nil, err
	} else if next { // in the white list, pass directly
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	userClaims, err := ParseAccessToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	ctx = logging.InjectFields(ctx, logging.Fields{"auth.sub", token})

	// WARNING: In production define your own type to avoid context collisions.
	return context.WithValue(ctx, tokenInfoKey, userClaims), nil
}
