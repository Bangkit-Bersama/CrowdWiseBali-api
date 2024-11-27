package v1

import "github.com/labstack/echo/v4"

func Response(c echo.Context, code int, i interface{}, e error) error {
	response := map[string]interface{}{
		"status": "success",
	}

	if i != nil {
		response["data"] = i
	}

	if e != nil {
		response["message"] = e.Error()
	}

	if code >= 400 && code <= 499 {
		response["status"] = "fail"
	} else if code >= 500 && code <= 599 {
		response["status"] = "error"
	}

	return c.JSON(code, response)
}
