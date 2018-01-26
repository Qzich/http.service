package http_service

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (action Response) SetAuthorizationHeader(authHeader string) error {
	action.ResponseWriter.Header().Set("Authorization", authHeader)

	return nil
}

func (action Response) SetHeader(header string, value string) error {
	action.ResponseWriter.Header().Set(header, value)

	return nil
}

func (action Response) SetResponse(statusCode int, model interface{}) error {
	action.ResponseWriter.WriteHeader(statusCode)

	switch {
	case statusCode == http.StatusMovedPermanently:
		action.ResponseWriter.Header().Set("Location", model.(string))

		return nil
	}

	return json.NewEncoder(action.ResponseWriter).Encode(model)
}

func (action Response) SetStatusCode(statusCode int) error {
	action.ResponseWriter.WriteHeader(statusCode)
	return nil
}

func (action Response) SetCookie(cookie *http.Cookie) error {
	http.SetCookie(action.ResponseWriter, cookie)
	return nil
}

func (action Response) DisableCache() error {
	action.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	action.ResponseWriter.Header().Set("Pragma", "no-cache")
	action.ResponseWriter.Header().Set("Expires", "0")

	return nil
}
