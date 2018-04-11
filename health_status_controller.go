package http_service

import (
	"net/http"
)

type HealthStatusController struct {
	Patients map[string]Pinger
}

func (controller *HealthStatusController) Methods() []string {
	return []string{http.MethodGet}
}

func (controller *HealthStatusController) Route() string {
	return "/_service/status"
}

func (controller *HealthStatusController) Action(responseWriter http.ResponseWriter, request *http.Request) error {

	healthStatus := healthChecker(controller.Patients).Status()

	responseWriter.WriteHeader(healthStatus)

	return nil
}
