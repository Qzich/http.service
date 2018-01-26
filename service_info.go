package http_service

type buildInfo struct {
	Date   string `json:"date"`
	Branch string `json:"branch"`
	Commit string `json:"commit"`
}

type serviceInfo struct {
	Build   *buildInfo         `json:"build"`
	DepList *HealthCheckerInfo `json:"dependency_list"`
}
