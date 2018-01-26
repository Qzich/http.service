package http_service

import "net/http"

type LoggerInterface interface {
	Info(infos ... interface{})
	Debug(warns ...interface{})
	Error(errors ...interface{})
}

type ControllerInterface interface {
	Methods() []string
	Route() string
	Action(http.ResponseWriter, *http.Request) error
}

type RouteHandlerInterface interface {
	Handle(http.ResponseWriter, *http.Request) error
}

type HttpServerInterface interface {
	Serve(address string) error
	AddRoute(methods []string, route string, handler RouteHandlerInterface)
	AddHandlerAdapter(handlerAdapter HandlerAdapter)
	NotFoundHandler(handler RouteHandlerInterface)
}
