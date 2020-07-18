package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"

	"github.com/ws-tobalobs/pkg/models"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	redis *redis.Client
}

// InitMiddleware intialize the middleware
func InitMiddleware(r *redis.Client) *GoMiddleware {
	return &GoMiddleware{
		redis: r,
	}
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT")
		// c.Response().Header().Set("Access-Control-Allow-Headers", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-Custom-Header, Upgrade-Insecure-Requests")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		// Access-Control-Allow-Credentials
		if c.Request().Method == "OPTIONS" {
			// c.Response().Header()
			return c.JSON(http.StatusOK, "ok")
		}

		return next(c)
	}
}

func (m *GoMiddleware) JwtAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		notAuth := []string{"/api/penyimpangan-kondisi-tambak", "/api/user/login", "/api/user/register", "/api/tambak/monitor", "/api/tambak/monitor-menyimpang", "/api/user/verify", "/api/user/forgot", "api/save-tunnel"}
		requestPath := c.Request().URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next(c)
				return nil
			}
		}

		var resp models.Responses
		tokenHeader := c.Request().Header.Get("Authorization")

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			resp.Message = "Missing auth token"
			resp.Status = models.StatusFailed
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, resp)
		}

		// splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		// if len(splitted) != 2 {
		// 	resp.Message = "Invalid/Malformed auth token"
		// 	resp.Status = models.StatusFailed
		// 	c.Response().Header().Set(`X-Cursor`, "header")
		// 	return c.JSON(http.StatusForbidden, resp)
		// }

		// tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("asdfghjkl"), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			resp.Message = "Malformed authentication token or token is expired"
			resp.Status = models.StatusFailed
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, resp)
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			resp.Message = "Token is not valid."
			resp.Status = models.StatusFailed
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, resp)
		}

		//Check token if token is stored in redis or InBlackList
		inBlackList := isInBlacklist(tokenHeader, m.redis)
		// log.Println(inBlackList)
		if inBlackList {
			resp.Message = "Token is not valid or InBlacklist"
			resp.Status = models.StatusFailed
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, resp)
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.UserId) //Useful for monitoring
		r := c.Request().WithContext(context.WithValue(c.Request().Context(), "user", tk.UserId))
		c.SetRequest(r)
		next(c) //proceed in the middleware chain!

		return nil
	})
}

func isInBlacklist(token string, redis *redis.Client) bool {
	res := redis.Get(token).Err()

	if res != nil {
		return false
	}

	return true
}
