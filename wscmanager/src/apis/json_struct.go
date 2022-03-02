package wsc_apis

type Loginjson struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Signupjson struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

type NotionCrawljson struct {
	Target    string `json:"target"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Type      string `json:"type"`
}

type SaveDBjson struct {
	Target    string `json:"target"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type WriteReportjson struct {
	Target    string `json:"target"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}
