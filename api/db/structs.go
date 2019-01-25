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
	Favorite  *int   `json:"favorite"`
	Comment   *int   `json:"comment"`
	Imagepath string `json:"imagepath"`
	Datetime  string `json:"datetime"`
}

type StructUser struct {
	Name     string `json:"name"`
	ShowName string `json:"showname"`
	Alive    bool   `json:"alive"`
}
