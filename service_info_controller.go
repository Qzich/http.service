package http_service

import (
	"encoding/json"
	"net/http"
)

type ServiceInfoController struct {
	Patients  map[string]Pinger
	BuildInfo *buildInfo
}

func (controller *ServiceInfoController) Methods() []string {
	return []string{http.MethodGet}
}

func (controller *ServiceInfoController) Route() string {
	return "/service/info"
}

func (controller *ServiceInfoController) Action(responseWriter http.ResponseWriter, request *http.Request) error {

	healthInfo := healthChecker(controller.Patients).Info()

	serviceInfo := &serviceInfo{
		Build:   controller.BuildInfo,
		DepList: &healthInfo,
	}

	return json.NewEncoder(responseWriter).Encode(serviceInfo)
}
