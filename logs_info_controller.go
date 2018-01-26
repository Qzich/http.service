package http_service

import (
	"io/ioutil"
	"net/http"
	"os"
)

type LogsInfoController struct {
	PathToFile string
}

func (*LogsInfoController) Methods() []string {
	return []string{http.MethodGet}
}

func (*LogsInfoController) Route() string {
	return "/dev/logs"
}

func (l *LogsInfoController) Action(responseWriter http.ResponseWriter, request *http.Request) error {

	SetDebugMode(false)

	file, _ := os.Open(l.PathToFile)
	logs, _ := ioutil.ReadAll(file)

	_, err := responseWriter.Write(logs)

	return err
}
