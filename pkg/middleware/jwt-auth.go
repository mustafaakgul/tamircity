package middleware

/*
import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	//"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	//"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
*/

/*
func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "Authorization header is missing", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Print("Claim[user_id]", claims["user_id"])
			log.Print("Claim[issuer]", claims["issuer"])
		} else {
			log.Print(err)
			response := helpers.BuildErrorResponse("Failed to process request", "Invalid token", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}

func ValidateJWT(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			render.Unauthorized(w, errors.New("authorization not found in header"))
			//http.Error(w, "authorization not found in header", http.StatusUnauthorized)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			render.Unauthorized(w, err)
			//http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "userAuthCtx", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "err.Error()", http.StatusUnauthorized)

	}
	return http.HandlerFunc(fn)
}
*/
