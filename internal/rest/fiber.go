package rest

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"reflect"
)

// JSONEncoder returns [] then slice or array is nil
func JSONEncoder(v interface{}) ([]byte, error) {
	if v == nil { // todo better solution ?
		typ := reflect.TypeOf(v).Kind()
		if typ == reflect.Array || typ == reflect.Slice {
			return []byte("[]"), nil
		}
	}
	return json.Marshal(v)
}

func errorResponse(c *fiber.Ctx, code int, logMessage error) error {
	log.Printf("handler error: %s", logMessage)
	// todo if debug .SendString(logMessage)
	c.Status(code)

	return nil
}
