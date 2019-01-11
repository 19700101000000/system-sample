package db

type OptionCategory struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type Gallery struct {
	Username  string `json:"username"`
	Imagepath string `json:"imagepath"`
	Datetime  string `json:"datetime"`
}
