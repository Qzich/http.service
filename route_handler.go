package http_service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

type responseRecorder struct {
	*httptest.ResponseRecorder

	wroteHeader bool
}

func (recorderResponse *responseRecorder) WriteHeader(code int) {
	recorderResponse.ResponseRecorder.WriteHeader(code)
	recorderResponse.wroteHeader = true
}

type RouteHandlerFunc func(http.ResponseWriter, *http.Request) error

func (routerHandler RouteHandlerFunc) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	routerHandler.Handle(response, request)
}

func (routerHandler RouteHandlerFunc) Handle(responseWriter http.ResponseWriter, request *http.Request) error {
	var err error

	responseRecorder := &responseRecorder{ResponseRecorder: httptest.NewRecorder()}
	responseRecorder.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	newContext := request.Context()

	for key, value := range vars {
		newContext = context.WithValue(newContext, "request."+key, value)
	}

	newRequest := request.WithContext(newContext)

	if err = routerHandler(responseRecorder, newRequest); err != nil {
		if !responseRecorder.wroteHeader {
			responseRecorder.Code = http.StatusInternalServerError
		}
	}

	for headerKey, headers := range responseRecorder.HeaderMap {
		responseWriter.Header().Del(headerKey)

		for _, header := range headers {
			responseWriter.Header().Add(headerKey, header)
		}
	}

	responseWriter.WriteHeader(responseRecorder.Code)
	responseWriter.Write(responseRecorder.Body.Bytes())

	defer func() {
		var dumpedRequest []byte

		if request.URL.Path == "/_service/info" {
			return
		}

		if request.URL.Path == "/_service/status" {
			return
		}

		dumpedRequest, err = httputil.DumpRequest(request, true)

		logger.Debug(
			"",
			"--- REQUEST ---: ",
			string(dumpedRequest),
			"",
			"--- RESPONSE ---: ",
			fmt.Sprintf("Status Code: %d", responseRecorder.Code),
			fmt.Sprintf("Headers: %v", responseRecorder.Header()),
			fmt.Sprintf("Body: %s", responseRecorder.Body.String()),
		)
	}()

	return err
}
