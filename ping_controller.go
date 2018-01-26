package http_service

import (
	"encoding/json"
	"net/http"
)

type PingController struct {
}

func (controller *PingController) Methods() []string {
	return []string{http.MethodGet}
}

func (controller *PingController) Route() string {
	return "/ping"
}

func (controller *PingController) Action(responseWriter http.ResponseWriter, request *http.Request) error {
	pingPongMap := map[string]string{"ping": "pong"}
	pingPongBytes, err := json.Marshal(pingPongMap)

	responseWriter.Write(pingPongBytes)

	return err
}
