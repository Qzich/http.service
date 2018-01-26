package http_service

import (
	"net/http"
)

type ControllersList []ControllerInterface

func (controllersList *ControllersList) Add(controllers ...ControllerInterface) {
	for _, controller := range controllers {
		*controllersList = append(*controllersList, controller)
	}
}

func CreateControllerHandler(controller ControllerInterface) RouteHandlerInterface {
	return &controllerHandler{controller: controller}
}

type controllerHandler struct {
	controller ControllerInterface
}

func (this *controllerHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) error {
	var err error

	if err = this.controller.Action(responseWriter, request); err != nil {
		logger.Debug("Action error: " + err.Error())
	}

	return err
}
