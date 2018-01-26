package http_service

import (
	"encoding/json"
	"net/http"
)

type notFoundHandler struct {
}

func (*notFoundHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) error {
	responseWriter.WriteHeader(http.StatusNotFound)

	responseBytes, err := json.Marshal(NotFoundResponse)

	_, err = responseWriter.Write(responseBytes)

	return err
}
