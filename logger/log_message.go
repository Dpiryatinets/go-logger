package logger

type logMessage struct {
	Message interface{} `json:"message"`
	Type string `json:"type"`
	ProcessId int `json:"processId"`
	Date string `json:"date"`
	ServiceName string `json:"serviceName"`
}
