package json

type trace struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Function string `json:"function"`
}
