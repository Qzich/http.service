package http_service

import (
	"net/http"
	"time"
)

type Pinger interface {
	Ping() (PingResponse, error)
}

type PingResponse struct {
	StatusCode int
	Latency    time.Duration
}

type pingResponse struct {
	StatusCode int     `json:"status"`
	Latency    float64 `json:"latency"`
}

type PingFunc func() (time.Duration, error)

func (f PingFunc) Ping() (PingResponse, error) {
	latency, err := f()

	if err != nil {
		return PingResponse{
			StatusCode: http.StatusBadRequest,
			Latency:    latency,
		}, err
	}

	return PingResponse{
		StatusCode: http.StatusOK,
		Latency:    latency,
	}, nil
}

type HealthCheckerInfo map[string]pingResponse

type healthStatus healthChecker

type healthChecker map[string]Pinger

func (healthChecker healthChecker) Status() int {

	for _, pinger := range healthChecker {
		if _, pingError := pinger.Ping(); pingError != nil {
			return http.StatusBadRequest
		}

	}

	return http.StatusOK
}

func (healthChecker healthChecker) Info() HealthCheckerInfo {

	healthCheckerInfo := HealthCheckerInfo{}

	for key, pinger := range healthChecker {

		response, _ := pinger.Ping()

		healthCheckerInfo[key] = pingResponse{
			StatusCode: response.StatusCode,
			Latency:    response.Latency.Seconds(),
		}

	}

	return healthCheckerInfo
}

func (healthChecker healthChecker) Add(key string, patient Pinger) healthChecker {
	healthChecker[key] = patient

	return healthChecker
}

func (healthChecker healthStatus) Add(patient Pinger) healthStatus {
	key := len(healthChecker)
	healthChecker[string(key)] = patient

	return healthChecker
}
