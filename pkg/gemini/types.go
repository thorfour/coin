package gemini

type Quote struct {
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	Last      string `json:"last"`
	TimeStamp uint64 `json:"timestamp"`
}
