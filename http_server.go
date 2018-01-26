package http_service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type routeHandlersGroup struct {
	routePattern string
	methods      []string
	handler      RouteHandlerInterface
}

func NewHttpServer() *httpServer {
	return &httpServer{
		server:          &http.Server{},
		routing:         []routeHandlersGroup{},
		notFoundHandler: &notFoundHandler{},
	}
}

type httpServer struct {
	server          *http.Server
	routing         []routeHandlersGroup
	handlerAdapters []HandlerAdapter
	notFoundHandler RouteHandlerInterface
}

func (server *httpServer) NotFoundHandler(handler RouteHandlerInterface) {
	server.notFoundHandler = handler
}

func (server *httpServer) Serve(address string) error {
	var httpHandler http.Handler

	serverMux := mux.NewRouter()
	httpHandler = serverMux

	for _, handlersGroup := range server.routing {
		serverMux.Handle(handlersGroup.routePattern, RouteHandlerFunc(handlersGroup.handler.Handle)).Methods(handlersGroup.methods...)
	}

	for _, handlerAdapter := range server.handlerAdapters {
		httpHandler = handlerAdapter(httpHandler)
	}

	serverMux.NotFoundHandler = RouteHandlerFunc(server.notFoundHandler.Handle)

	server.server.Addr = address
	server.server.Handler = httpHandler

	return server.server.ListenAndServe()
}

func (server *httpServer) AddRoute(methods []string, route string, handler RouteHandlerInterface) {
	server.routing = append(server.routing, routeHandlersGroup{routePattern: route, methods: methods, handler: handler})
}

func (server *httpServer) AddHandlerAdapter(handlerAdapter HandlerAdapter) {
	server.handlerAdapters = append(server.handlerAdapters, handlerAdapter)
}
