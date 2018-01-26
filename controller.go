package http_service

import "net/http"

func NewController(methods []string, route string, action func(http.ResponseWriter, *http.Request) error) *Controller {
	return &Controller{methods: methods, route: route, action: action}
}

type Controller struct {
	methods []string
	route   string
	action  func(http.ResponseWriter, *http.Request) error
}

func (controller *Controller) Methods() []string {
	return controller.methods
}

func (controller *Controller) Route() string {
	return controller.route
}

func (controller *Controller) Action(responseWriter http.ResponseWriter, request *http.Request) error {
	return controller.action(responseWriter, request)
}
