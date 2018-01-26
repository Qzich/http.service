package http_service

var (
	HandlerAdapters = HandlerAdaptersList{}
	Controllers     = ControllersList{}
	ServiceInfo     = healthChecker{}
	BuildInfo       = buildInfo{}
	HealthStatus    = healthStatus{}
	HttpServer      HttpServerInterface
	ServerAddress   string
	RoutePrefix     string
	debugMode       bool
	logger          LoggerInterface
)

func init() {
	logger = NewLogger(debugMode)
	HttpServer = NewHttpServer()
	ServerAddress = ":8080"
}

func Run() error {
	logger.Info("setup http service")

	healthInfoController, healthStatusController :=
		&ServiceInfoController{Patients: ServiceInfo, BuildInfo: &BuildInfo},
		&HealthStatusController{HealthStatus}

	controllers :=
		Controllers

	if debugMode {
		controllers.Add(&PingController{})
	}

	controllers.Add(healthInfoController)
	controllers.Add(healthStatusController)

	for _, controller := range controllers {
		HttpServer.AddRoute(
			controller.Methods(),
			RoutePrefix+controller.Route(),
			CreateControllerHandler(controller))
	}

	for _, handlerAdapter := range HandlerAdapters {
		HttpServer.AddHandlerAdapter(handlerAdapter)
	}

	logger.Info("http service listen on " + ServerAddress)

	return HttpServer.Serve(ServerAddress)
}

func SetDebugMode(flag bool) {
	debugMode = flag

	logger = NewLogger(flag)
}

func Logger() LoggerInterface {
	return logger
}
