package helper

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/phungvandat/life-cafe-backend/util/config"
	"github.com/phungvandat/life-cafe-backend/util/contextkey"
)

func VerifyToken(ctx context.Context, req *http.Request) context.Context {
	accessToken := req.Header.Get("Authorization")
	if strings.Trim(accessToken, " ") != "" {
		claims, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, InvalidSigningAlgorithm
			}
			secret := config.GetJWTSerectKeyEnv()
			return []byte(secret), nil
		})

		if err != nil || !claims.Valid {
			goto End
		}
		data := claims.Claims.(jwt.MapClaims)
		userID, check := data["user_id"].(string)

		if check {
			ctx = context.WithValue(ctx, contextkey.UserIDContextKey, userID)
		}

		username, check := data["username"].(string)
		if check {
			ctx = context.WithValue(ctx, contextkey.UsernameContextKey, username)
		}

		role, check := data["role"].(string)
		if check {
			ctx = context.WithValue(ctx, contextkey.UserRoleContextKey, role)
		}
	}
End:
	return ctx
}
