package common

// Msg format
type Msg struct {
	Time    string `json:"time"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
