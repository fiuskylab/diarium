package json

type LogJSON struct {
	Level       string      `json:"level"`
	Message     string      `json:"message"`
	Trace       trace       `json:"trace"`
	Network     network     `json:"network"`
	Application application `json:"application"`
}
