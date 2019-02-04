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

type StructWanted struct {
	Username    string `json:"username"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Alive       bool   `json:"alive"`
	RequestNum  *int   `json:"requests"`
}

type StructRequest struct {
	Wanted      StructWanted `json:"wanted"`
	Number      int          `json:"number"`
	UserName    string       `json:"username"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Price       int          `json:"price"`
	Establish   bool         `json:"establish"`
	Alive       bool         `json:"alive"`
	Check       bool         `json:"check"`
}

type StructInfo struct {
	WantedNum  int `json:"wanteds"`
	RequestNum int `json:"requests"`
}
