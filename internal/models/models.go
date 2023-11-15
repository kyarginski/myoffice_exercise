package models

type Report struct {
	URL   string `json:"url"`
	Size  int64  `json:"size"`
	Time  string `json:"time"`
	Error string `json:"error"`
}
