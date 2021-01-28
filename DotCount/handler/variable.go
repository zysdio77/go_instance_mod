package handler

type UserDetailJson struct {
	LogoutTime   int64  `json:"logoutTime" form:"logoutTime"`
	DevModel     string `json:"devModel" form:"devModel"`
	Platform     string `json:"platform" form:"platform"`
	Channel      string `json:"channel" form:"channel"`
	WildStamp    string `json:"wildStamp" form:"wildStamp"`
	UserGambling string `json:"userGambling" form:"userGambling"`
	TotalWin     string `json:"totalWin" form:"totalWin"`
	TotalBet     string `json:"totalBet" form:"totalBet"`
	UserValue    string `json:"userValue" form:"userValue"`
	FreeCoin     string `json:"freeCoin" form:"freeCoin"`
	Version      string `json:"version" form:"version"`
	UID          int    `json:"uid" form:"uid,string"`
	CoinNum      string `json:"coinNum" form:"coinNum"`
	Vip          int    `json:"vip" form:"vip"`
	Lv           int    `json:"lv" form:"lv"`
	GuideF       string `json:"guideF" form:"guideF"`
	GuideNF      string `json:"guideNF" form:"guideNF"`
}

type UserDetailJsonRaw struct {
	Country      string `json:"country" form:"country"`
	Lang         string `json:"lang" form:"lang"`
	Time         int64  `json:"time" form:"time"`
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

type FreeCoinDataJson struct {
	Uid             int    `json:"uid" form:"uid"`
	Channel         string `json:"channel" form:"channel"`
	Platform        string `json:"platform" form:"platform"`
	FreeCoinDetails string `json:"freeCoinDetails" form:"freeCoinDetails"`
	Version         string `json:"version" form:"version"`
}

type FreeCoinDetailJson []struct {
	TimeStamp int64    `json:"timeStamp"`
	CoinNum   int    `json:"coinNum"`
	CoinName  string `json:"coinName"`
}
type FreeCoinDetailJson2 struct {
	TimeStamp int64    `json:"timeStamp"`
	CoinNum   int    `json:"coinNum"`
	CoinName  string `json:"coinName"`
}


type SlotDetailJson struct {
	//Country     string `json:"country"`
	//Lang        string `json:"lang"`
	Time        int64 `json:"time" form:"time"`
	//Udid        string `json:"udid"`
	//Adid        string `json:"adid"`
	//UUID        string `json:"uuid"`
	//Idfv        string `json:"idfv"`
	//Hmac        string `json:"hmac"`
	//Email       string `json:"email"`
	//DevModel    string `json:"devModel"`
	//Brand       string `json:"brand"`
	IsTestUser  int    `json:"isTestUser" form:"isTestUser"`
	ClientVer   string `json:"clientVer" form:"clientVer"`
	ClientBuild string `json:"clientBuild" form:"clientBuild"`
	OsType      string `json:"osType" form:"osType"`
	OsVer       string `json:"osVer" form:"osVer"`
	PackageName string `json:"packageName" form:"packageName"`
	Platform    string `json:"platform" form:"platform"`
	Channel     string `json:"channel" form:"channel"`
	UID         int    `json:"uid" form:"uid"`
	SlotDetails string `json:"slotDetails" form:"slotDetails"`
	Version     string    `json:"version" form:"version"`
}

type SlotDetails struct {
	WinNum int    `json:"winNum"`
	BetNum int    `json:"betNum"`
	SlotID string `json:"slotId"`
}

type PayResultJson struct {
	SkuID     string `json:"skuID"`
	Price     int    `json:"price"`
	PayKey    string `json:"payKey"`
	PayCoupon int    `json:"payCoupon"`
	PayDetail string `json:"payDetail"`
	Tid       string `json:"tid"`
	UID       int    `json:"uid"`
}
type PayDetails struct {
	Coin      int64 `json:"coin"`
	RepeatWin struct {
		ID     string `json:"id"`
		From   string `json:"from"`
		IapKey string `json:"iapKey"`
	} `json:"RepeatWin"`
	StoreBlast struct {
		ID     string `json:"id"`
		IapKey string `json:"iapKey"`
	} `json:"StoreBlast"`
	Picks     int      `json:"picks"`
	Cn        []string `json:"cn"`
	Vp        int      `json:"vp"`
	Sr        int      `json:"sr"`
	PayMoney  int      `json:"payMoney"`
	CoinStart int64    `json:"coinStart"`
	CoinEnd   int64    `json:"coinEnd"`
	StampNum  string   `json:"stamp_num"`
	Postmark  struct {
		Mark int `json:"mark"`
		Coin int `json:"coin"`
	} `json:"postmark"`
}