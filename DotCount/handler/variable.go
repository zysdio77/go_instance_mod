package handler

type UserDetailJson struct {
	Country      string `json:"country" form:"country"`
	Lang         string `json:"lang" form:"lang"`
	Time         int64 `json:"time" form:"time"`
	Udid         string `json:"udid" form:"udid"`
	Adid         string `json:"adid" form:"adid"`
	UUID         string `json:"uuid" form:"uuid"`
	Idfv         string `json:"idfv" form:"idfv"`
	Hmac         string `json:"hmac" form:"hmac"`
	Email        string `json:"email" form:"email"`
	DevModel     string `json:"devModel" form:"devModel"`
	Brand        string `json:"brand" form:"brand"`
	IsTestUser   int    `json:"isTestUser" form:"isTestUser"`
	ClientVer    string `json:"clientVer" form:"clientVer"`
	ClientBuild  int    `json:"clientBuild" form:"clientBuild"`
	OsType       int    `json:"osType" form:"osType"`
	OsVer        int    `json:"osVer" form:"osVer"`
	PackageName  string `json:"packageName" form:"packageName"`
	Platform     string `json:"platform" form:"platform"`
	Channel      string `json:"channel" form:"channel"`
	WildStamp    string `json:"wildStamp" form:"wildStamp"`
	UserGambling string `json:"userGambling" form:"userGambling"`
	TotalWin     string `json:"totalWin" form:"totalWin"`
	TotalBet     string `json:"totalBet" form:"totalBet"`
	UserValue    string `json:"userValue" form:"userValue"`
	FreeCoin     string `json:"freeCoin" form:"freeCoin"`
	Version      string `json:"version" form:"version"`
	UID          int    `json:"uid" form:"uid"`
	CoinNum      string `json:"coinNum" form:"coinNum"`
	Vip          int    `json:"vip" form:"vip"`
	Lv           int    `json:"lv" form:"lv"`
	GuideF       string `json:"guideF" form:"guideF"`
	GuideNF      string `json:"guideNF" form:"guideNF"`
}


type FreeCoinDetail struct {
}
type c struct {
}
type d struct {
}
type e struct {
}
type f struct {
}
type g struct {
}
