package bing

// WordModel 单词模型
type WordModel struct {
	Word          string `json:"word"`
	Pronunciation struct {
		AmE    string `json:"AmE"`
		AmEmp3 string `json:"AmEmp3"`
		BrE    string `json:"BrE"`
		BrEmp3 string `json:"BrEmp3"`
	} `json:"pronunciation"`
	Defs []struct {
		Pos string `json:"pos"`
		Def string `json:"def"`
	} `json:"defs"`
	Sams []struct {
		Eng    string `json:"eng"`
		Chn    string `json:"chn"`
		Mp3Url string `json:"mp3Url"`
	} `json:"sams"`
}
