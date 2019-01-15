package db

type OptionCategory struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type Gallery struct {
	Username  string `json:"username"`
	Number    int    `json:"number"`
	EvalSum   *int   `json:"evalsum"`
	EvalParam *int   `json:"evalparam"`
	Imagepath string `json:"imagepath"`
	Datetime  string `json:"datetime"`
}
