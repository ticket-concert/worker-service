package middleware

import (
	"fmt"
	"strings"
	config "worker-service/configs"
	"worker-service/internal/pkg/constants"
	"worker-service/internal/pkg/errors"
	helpers "worker-service/internal/pkg/helpers"
	"worker-service/internal/pkg/log"
	"worker-service/internal/pkg/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type Middlewares struct {
	redisClient redis.Collections
}

func NewMiddlewares(redis redis.Collections) Middlewares {
	return Middlewares{
		redisClient: redis,
	}
}

func (m Middlewares) VerifyBasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.GetConfig().UsernameBasicAuth: config.GetConfig().PasswordBasicAuth,
		},
	})
}

func (m Middlewares) VerifyBearer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := log.GetLogger()
		redisClient := m.redisClient
		helperImpl := &helpers.JwtImpl{}
		parseToken, err := helperImpl.JWTAuthorization(c.Request())
		if err != nil {
			return helpers.RespError(c, logger, err)
		}
		token := strings.Split(string(c.Request().Header.Peek("Authorization")), " ")
		if len(token) != 2 || (token[0] != "Bearer" && token[0] != "bearer") {
			logger.Error(c.Context(), "Invalid token format", token[1])
			return helpers.RespError(c, logger, errors.ForbiddenError("Invalid token format"))
		}

		blocklist, _ := redisClient.Get(c.Context(), fmt.Sprintf("%s:%s", constants.RedisKeyBlockListJwt, token[1])).Result()
		if blocklist != "" {
			logger.Error(c.Context(), "Access token expired!", "Token blocklist")
			return helpers.RespError(c, logger, errors.UnauthorizedError("Access token expired!"))
		}
		c.Locals("userId", parseToken.UserId)
		c.Locals("userRole", parseToken.Role)
		return c.Next()
	}

}
