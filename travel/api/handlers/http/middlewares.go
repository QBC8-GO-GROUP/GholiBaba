package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TraceMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		traceID := c.Get("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c.Set("X-Trace-ID", traceID)

		c.Locals("traceID", traceID)

		return c.Next()
	}
}

func method2ScopeMapper(method string) string {
	if method == "GET" {
		return "read"
	}

	switch method {
	case "GET":
		return "read"
	case "POST":
		return "create"
	case "HEAD":
		return "create"
	case "DELETE":
		return "delete"
	case "PUT":
		return "update"
	case "PATCH":
		return "patch"
	default:
		return ""
	}
}
